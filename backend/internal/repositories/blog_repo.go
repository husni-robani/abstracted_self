package repositories

import (
	"database/sql"
	"fmt"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(DB *sql.DB) BlogRepository{
	return BlogRepository{
		db: DB,
	}
}

func (repo BlogRepository) GetAllBlogs() ([]models.Blog, error){
	var blogs []models.Blog
	rows, err := repo.db.Query("Select id, title, url, image, blog_snippet from blogs")
	if err != nil {
		logger.Error.Printf("get blog data error: %#v", err.Error())
		return nil, err
	}

	for rows.Next(){
		blog := models.Blog{}

		err := rows.Scan(&blog.Id, &blog.Title, &blog.URL, &blog.Image, &blog.BlogSnippet)
		if err != nil {
			logger.Error.Printf("scan rows error: %#v", err.Error())
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repo BlogRepository) GetBlogByID(id int) (models.Blog, error) {
	blog := models.Blog{}

	row := repo.db.QueryRow("select * from blogs where id = $1", id)

	err := row.Scan(&blog.Id, &blog.Title, &blog.URL, &blog.Image)
	if err != nil {
		logger.Error.Printf("error select blog: %#v", err.Error())
		return blog, err
	}

	return blog, nil
}

func (repo BlogRepository) CreateBlog(blog requests.CreateBlogRequest) error {
	result, err := repo.db.Exec("INSERT INTO blogs (title, url, image, blog_snippet) VALUES ($1, $2, $3, $4)", blog.Title, blog.URL, blog.Image, blog.BlogSnippet)
	if err != nil {
		logger.Error.Printf("error insert blog: %#v", err.Error())
		return err
	}
	logger.Info.Print("blog created!")

	totalRowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("error get total rows affected: %#v", err.Error())
		return err
	}

	logger.Info.Printf("row affected: %v", totalRowsAffected)

	return nil
}

func (repo BlogRepository) DeleteBlog(id int) error {
	result, err := repo.db.Exec("DELETE FROM blogs where id = $1", id)
	if err != nil {
		 logger.Error.Printf("error delete blog: %v", err.Error())
		 return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("error get total rows affected: %v", err.Error())
		return err
	}

	logger.Info.Printf("row affected: %v", rowsAffected)
	
	return nil
}

func (repo BlogRepository) UpdateBlog(id int, updateValues map[string]any) error {
	query := utils.GenerateSingleUpdateQuery("blogs", updateValues, fmt.Sprintf("where id = %d", id))

	result, err := repo.db.Exec(query)
	if err != nil {
		logger.Error.Printf("error exec query update: %s\n query: %s", err.Error(), query)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("error get row affected: %s", err.Error())
		return err
	}

	logger.Info.Printf("Row affected: %v", rowsAffected)

	return nil
}