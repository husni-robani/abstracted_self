package services

import (
	"mime/multipart"
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
	ImageRepo  repositories.ImageRepository
}

func NewBlogService(blogRepo repositories.BlogRepository, imageRepo repositories.ImageRepository) BlogService {
	return BlogService{Repository: blogRepo, ImageRepo: imageRepo}
}

func (service BlogService) GetAllBlogs(published *bool) ([]models.Blog, error) {
	blogs, err := service.Repository.GetAllBlogs(published)
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

	contentImages, err := service.Repository.GetContentImages(id)
	if err != nil {
		return models.Blog{}, err
	}

	blog.Images = contentImages

	return blog, nil
}

func (service BlogService) saveCoverImage(fileHeader *multipart.FileHeader) (models.Image, error) {
	extension := filepath.Ext(fileHeader.Filename)
	newFilename := uuid.New().String() + extension
	mimeType := fileHeader.Header.Get("Content-Type")

	fileHeader.Filename = newFilename
	if err := utils.SaveFile(fileHeader, "."+os.Getenv("IMAGES_STORAGE_PATH")); err != nil {
		return models.Image{}, err
	}

	imageId, err := service.ImageRepo.CreateImage(newFilename, fileHeader.Size, mimeType)
	if err != nil {
		return models.Image{}, err
	}

	return models.Image{
		Id:       imageId,
		URL:      models.ImageURL(imageId),
		FileName: newFilename,
		FileSize: int(fileHeader.Size),
		MimeType: mimeType,
	}, nil
}

func (service BlogService) CreateBlog(blogData requests.CreateBlogRequest) error {
	if err := service.ImageRepo.ValidateImagesExist(blogData.ContentImageIds); err != nil {
		return err
	}

	coverImage, err := service.saveCoverImage(blogData.ImageFile)
	if err != nil {
		return err
	}

	blog := models.Blog{
		Title:        blogData.Title,
		Slug:         blogData.Slug,
		Content:      blogData.Content,
		BlogSnippet:  blogData.BlogSnippet,
		Published:    blogData.Published,
		CoverImageId: coverImage.Id,
	}

	if blog.Slug == "" {
		blog.Slug = generateSlug(blog.Title)
	}

	if err := service.Repository.CreateBlog(blog, blogData.ContentImageIds); err != nil {
		return err
	}

	return nil
}

func (service BlogService) DeleteBlog(id int) error {
	blog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		return err
	}

	contentImages, err := service.Repository.GetContentImages(id)
	if err != nil {
		return err
	}

	if err := service.Repository.DeleteBlog(id); err != nil {
		return err
	}

	var firstErr error
	if blog.CoverImage != nil {
		if err := service.ImageRepo.DeleteImage(blog.CoverImage.Id); err != nil {
			logger.Error.Printf("error delete cover image: %v", err.Error())
			firstErr = err
		} else if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", blog.CoverImage.FileName); err != nil {
			logger.Error.Printf("error delete cover image file: %v", err.Error())
			firstErr = err
		}
	}

	for _, image := range contentImages {
		if err := service.ImageRepo.DeleteImage(image.Id); err != nil {
			logger.Error.Printf("error delete content image: %v", err.Error())
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", image.FileName); err != nil {
			logger.Error.Printf("error delete content image file: %v", err.Error())
			if firstErr == nil {
				firstErr = err
			}
		}
	}

	return firstErr
}

func (service BlogService) UpdateBlog(id int, newBlogData requests.UpdateBlogRequest) (*models.Blog, error) {
	if err := service.ImageRepo.ValidateImagesExist(newBlogData.ContentImageIds); err != nil {
		return nil, err
	}

	blog, err := service.Repository.GetBlogByID(id)
	if err != nil {
		return nil, err
	}

	oldContentImages, err := service.Repository.GetContentImages(id)
	if err != nil {
		return nil, err
	}

	coverImageId := blog.CoverImageId
	var oldCoverImage *models.Image

	if newBlogData.ImageFile != nil {
		coverImage, err := service.saveCoverImage(newBlogData.ImageFile)
		if err != nil {
			logger.Error.Printf("error save cover image: %v", err.Error())
			return nil, err
		}
		oldCoverImage = blog.CoverImage
		coverImageId = coverImage.Id
	}

	blog.Title = newBlogData.Title
	blog.Content = newBlogData.Content
	blog.BlogSnippet = newBlogData.BlogSnippet
	blog.Published = newBlogData.Published
	blog.CoverImageId = coverImageId

	if newBlogData.Slug != "" {
		blog.Slug = newBlogData.Slug
	} else {
		blog.Slug = generateSlug(newBlogData.Title)
	}

	now := time.Now()
	blog.UpdatedAt = &now

	if err := service.Repository.UpdateBlog(id, blog, newBlogData.ContentImageIds); err != nil {
		return nil, err
	}

	if oldCoverImage != nil {
		if err := service.ImageRepo.DeleteImage(oldCoverImage.Id); err != nil {
			logger.Error.Printf("error delete old cover image: %v", err.Error())
		} else if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", oldCoverImage.FileName); err != nil {
			logger.Error.Printf("error delete old cover image file: %v", err.Error())
		}
	}

	newImageIds := make(map[int]bool, len(newBlogData.ContentImageIds))
	for _, imageId := range newBlogData.ContentImageIds {
		newImageIds[imageId] = true
	}

	for _, image := range oldContentImages {
		if newImageIds[image.Id] {
			continue
		}

		if err := service.ImageRepo.DeleteImage(image.Id); err != nil {
			logger.Error.Printf("error delete removed content image: %v", err.Error())
			continue
		}
		if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", image.FileName); err != nil {
			logger.Error.Printf("error delete removed content image file: %v", err.Error())
		}
	}

	updatedBlog, err := service.GetBlogByID(id)
	if err != nil {
		return nil, err
	}

	return &updatedBlog, nil
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
