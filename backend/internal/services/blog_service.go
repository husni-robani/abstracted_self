package services

import (
	"path/filepath"

	"github.com/google/uuid"
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