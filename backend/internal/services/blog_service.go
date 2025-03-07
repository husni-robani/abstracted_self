package services

import (
	"os"
	"path/filepath"

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

func (service BlogService) CreateBlog(blog models.Blog) error {
	// change filename
	extension := filepath.Ext(blog.ImageFile.Filename)
	newFilename := uuid.New().String() + extension

	blog.ImageFile.Filename = newFilename
	blog.Image = newFilename

	// save image
	err := utils.SaveFile(blog.ImageFile, "./storage/blog_images")
	if err != nil {
		return err
	}

	// save to database
	err = service.Repository.CreateBlog(blog)
	if err != nil {
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

	if err := os.Remove("./storage/blog_images/" + blog.Image); err != nil {
		logger.Error.Printf("error delete image: %v", err.Error())
		return err
	}

	return nil
}

func (service BlogService) UpdateBlog(id int, newBlogData requests.UpdateBlogRequest) (*models.Blog, error) {
	blog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		logger.Error.Printf("error while get blog data: %v", err)
		return nil, err
	}

	var values = make(map[string]any)

	if newBlogData.ImageFile != nil {
		// save image
		extension := filepath.Ext(newBlogData.ImageFile.Filename)
		newFilename := uuid.New().String() + extension
		
		newBlogData.ImageFile.Filename = newFilename
		if err := utils.SaveFile(newBlogData.ImageFile, "./storage/blog_images"); err != nil {
			logger.Error.Printf("error save image to storage: %v", err.Error())
			return nil, err
		}
		logger.Info.Printf("file saved: %v", newBlogData.ImageFile.Filename)
		
		// remove older image
		err := os.Remove("./storage/blog_images/" + blog.Image)
		if err != nil {
			logger.Error.Printf("error remove image from storage: %v", err.Error())
			return nil, err
		}
		values["image"] = newFilename
	}

	// save data to database
	if newBlogData.Title != nil {
		values["title"] = *newBlogData.Title
	} 
	if newBlogData.URL != nil {
		values["url"] = *newBlogData.URL
	} 

	if err := service.Repository.UpdateBlog(id, values); err != nil {
		return nil, err
	}

	updatedBlog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		return nil, err
	}

	return &updatedBlog, nil
}