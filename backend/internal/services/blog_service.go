package services

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type BlogService struct {
	Repository repositories.BlogRepository
}

func NewBlogService(blogRepo repositories.BlogRepository) BlogService {
	return BlogService{Repository: blogRepo}
}

func (service BlogService) GetAllBlogs() ([]models.Blog, error) {
	blogs, err := service.Repository.GetAllBlogs()
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (service BlogService) GetBlogByID(id int) (models.Blog, error) {
	blog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}

func (service BlogService) CreateBlog(blog requests.CreateBlogRequest) error {
	// save cover image
	if blog.ImageFile != nil {
		extension := filepath.Ext(blog.ImageFile.Filename)
		newFilename := uuid.New().String() + extension

		blog.ImageFile.Filename = newFilename
		blog.Image = newFilename

		if err := utils.SaveFile(blog.ImageFile, "."+os.Getenv("IMAGES_STORAGE_PATH")); err != nil {
			return err
		}
	}

	// generate slug from title if not provided
	if blog.Slug == "" {
		blog.Slug = generateSlug(blog.Title)
	}

	// save to database
	if err := service.Repository.CreateBlog(blog); err != nil {
		return err
	}

	return nil
}

func (service BlogService) DeleteBlog(id int) error {
	blog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		return err
	}

	if err := service.Repository.DeleteBlog(blog.Id); err != nil {
		return err
	}

	if blog.Image != "" {
		if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", blog.Image); err != nil {
			logger.Error.Printf("error delete image: %v", err.Error())
			return err
		}
	}

	return nil
}

func (service BlogService) UpdateBlog(id int, newBlogData requests.UpdateBlogRequest) (*models.Blog, error) {
	blog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		return nil, err
	}

	// save new cover image and remove the old one
	if newBlogData.ImageFile != nil {
		extension := filepath.Ext(newBlogData.ImageFile.Filename)
		newFilename := uuid.New().String() + extension

		newBlogData.ImageFile.Filename = newFilename
		if err := utils.SaveFile(newBlogData.ImageFile, "."+os.Getenv("IMAGES_STORAGE_PATH")); err != nil {
			logger.Error.Printf("error save image to storage: %v", err.Error())
			return nil, err
		}
		logger.Info.Printf("file saved: %v", newBlogData.ImageFile.Filename)

		if blog.Image != "" {
			if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", blog.Image); err != nil {
				logger.Error.Printf("error remove image from storage: %v", err.Error())
			}
		}
		blog.Image = newFilename
	}

	blog.Title = newBlogData.Title
	blog.Content = newBlogData.Content
	blog.BlogSnippet = newBlogData.BlogSnippet
	blog.Published = newBlogData.Published

	if newBlogData.Slug != "" {
		blog.Slug = newBlogData.Slug
	} else {
		blog.Slug = generateSlug(newBlogData.Title)
	}

	now := time.Now()
	blog.UpdatedAt = &now

	if err := service.Repository.UpdateBlog(id, blog); err != nil {
		return nil, err
	}

	return &blog, nil
}

func generateSlug(title string) string {
	var b strings.Builder
	lastHyphen := false

	for _, r := range strings.ToLower(title) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			lastHyphen = false
		} else if !lastHyphen && b.Len() > 0 {
			b.WriteByte('-')
			lastHyphen = true
		}
	}

	return strings.Trim(b.String(), "-")
}
