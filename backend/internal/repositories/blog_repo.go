package repositories

import (
	"database/sql"
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(DB *sql.DB) BlogRepository {
	return BlogRepository{
		db: DB,
	}
}

type rowScanner interface {
	Scan(dest ...any) error
}

const blogColumns = "id, title, slug, image, content, blog_snippet, published, created_at, updated_at"

func scanBlog(s rowScanner) (models.Blog, error) {
	blog := models.Blog{}

	var image sql.NullString
	var blogSnippet sql.NullString
	var updatedAt sql.NullTime

	err := s.Scan(
		&blog.Id,
		&blog.Title,
		&blog.Slug,
		&image,
		&blog.Content,
		&blogSnippet,
		&blog.Published,
		&blog.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return blog, err
	}

	if image.Valid {
		blog.Image = image.String
	}
	if blogSnippet.Valid {
		blog.BlogSnippet = blogSnippet.String
	}
	if updatedAt.Valid {
		blog.UpdatedAt = &updatedAt.Time
	}

	return blog, nil
}

func nullableString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func nullableTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func (repo BlogRepository) GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog

	rows, err := repo.db.Query("SELECT " + blogColumns + " FROM blogs ORDER BY id")
	if err != nil {
		logger.Error.Printf("get blog data error: %#v", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		blog, err := scanBlog(rows)
		if err != nil {
			logger.Error.Printf("scan rows error: %#v", err.Error())
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repo BlogRepository) GetBlogByID(id int) (models.Blog, error) {
	row := repo.db.QueryRow("SELECT "+blogColumns+" FROM blogs WHERE id = $1", id)

	blog, err := scanBlog(row)
	if err != nil {
		logger.Error.Printf("error select blog: %#v", err.Error())
		return blog, err
	}

	return blog, nil
}

func (repo BlogRepository) CreateBlog(blog requests.CreateBlogRequest) error {
	_, err := repo.db.Exec(
		"INSERT INTO blogs (title, slug, image, content, blog_snippet, published) VALUES ($1, $2, $3, $4, $5, $6)",
		blog.Title, blog.Slug, nullableString(blog.Image), blog.Content, nullableString(blog.BlogSnippet), blog.Published,
	)
	if err != nil {
		logger.Error.Printf("error insert blog: %#v", err.Error())
		return err
	}
	logger.Info.Print("blog created!")

	return nil
}

func (repo BlogRepository) DeleteBlog(id int) error {
	_, err := repo.db.Exec("DELETE FROM blogs WHERE id = $1", id)
	if err != nil {
		logger.Error.Printf("error delete blog: %v", err.Error())
		return err
	}

	logger.Info.Printf("blog deleted: %v", id)

	return nil
}

func (repo BlogRepository) UpdateBlog(id int, blog models.Blog) error {
	tx, err := repo.db.Begin()
	if err != nil {
		logger.Error.Printf("error begin transaction: %v", err.Error())
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec("DELETE FROM blogs WHERE id = $1", id); err != nil {
		logger.Error.Printf("error delete blog in update: %v", err.Error())
		return err
	}

	if _, err := tx.Exec(
		"INSERT INTO blogs (id, title, slug, image, content, blog_snippet, published, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		id, blog.Title, blog.Slug, nullableString(blog.Image), blog.Content, nullableString(blog.BlogSnippet), blog.Published, blog.CreatedAt, nullableTime(blog.UpdatedAt),
	); err != nil {
		logger.Error.Printf("error insert blog in update: %v", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		logger.Error.Printf("error commit transaction: %v", err.Error())
		return err
	}

	logger.Info.Printf("blog updated: %v", id)

	return nil
}
