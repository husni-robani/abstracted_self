package repositories

import (
	"database/sql"
	"time"

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

func nullableString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func nullableTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

const blogSelect = "SELECT b.id, b.title, b.slug, b.cover_image_id, i.file_name, i.file_size, i.mime_type, b.blog_snippet, b.published, b.created_at, b.updated_at"

func scanBlog(s rowScanner, blog *models.Blog, withContent bool) error {
	var coverImageId sql.NullInt32
	var coverFileName sql.NullString
	var coverFileSize sql.NullInt32
	var coverMimeType sql.NullString
	var updatedAt sql.NullTime
	var content sql.NullString

	dest := []any{
		&blog.Id,
		&blog.Title,
		&blog.Slug,
		&coverImageId,
		&coverFileName,
		&coverFileSize,
		&coverMimeType,
		&blog.BlogSnippet,
		&blog.Published,
		&blog.CreatedAt,
		&updatedAt,
	}
	if withContent {
		dest = append(dest, &content)
	}

	if err := s.Scan(dest...); err != nil {
		return err
	}

	blog.CoverImageId = int(coverImageId.Int32)
	if coverImageId.Valid {
		blog.CoverImage = &models.Image{
			Id:       int(coverImageId.Int32),
			URL:      models.ImageURL(int(coverImageId.Int32)),
			FileName: coverFileName.String,
			FileSize: int(coverFileSize.Int32),
			MimeType: coverMimeType.String,
		}
	}

	if updatedAt.Valid {
		blog.UpdatedAt = &updatedAt.Time
	}
	if withContent {
		blog.Content = content.String
	}

	return nil
}

func (repo BlogRepository) GetAllBlogs(published *bool) ([]models.Blog, error) {
	var blogs []models.Blog

	query := blogSelect + " FROM blogs b LEFT JOIN images i ON i.id = b.cover_image_id"
	if published != nil {
		query += " WHERE b.published = $1"
	}
	query += " ORDER BY b.id"

	var rows *sql.Rows
	var err error

	if published != nil {
		rows, err = repo.db.Query(query, *published)
	} else {
		rows, err = repo.db.Query(query)
	}
	if err != nil {
		logger.Error.Printf("get blog data error: %#v", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		blog := models.Blog{}
		if err := scanBlog(rows, &blog, false); err != nil {
			logger.Error.Printf("scan rows error: %#v", err.Error())
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repo BlogRepository) GetBlogByID(id int) (models.Blog, error) {
	blog := models.Blog{}

	row := repo.db.QueryRow(blogSelect+", b.content FROM blogs b LEFT JOIN images i ON i.id = b.cover_image_id WHERE b.id = $1", id)
	if err := scanBlog(row, &blog, true); err != nil {
		logger.Error.Printf("error select blog: %#v", err.Error())
		return blog, err
	}

	return blog, nil
}

func (repo BlogRepository) GetContentImages(blogId int) ([]models.Image, error) {
	rows, err := repo.db.Query(
		"SELECT i.id, i.file_name, i.file_size, i.mime_type FROM blog_images bi JOIN images i ON i.id = bi.image_id WHERE bi.blog_id = $1 ORDER BY i.id",
		blogId,
	)
	if err != nil {
		logger.Error.Printf("error select blog content images: %#v", err.Error())
		return nil, err
	}
	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		image := models.Image{}
		if err := rows.Scan(&image.Id, &image.FileName, &image.FileSize, &image.MimeType); err != nil {
			logger.Error.Printf("scan rows error: %#v", err.Error())
			return nil, err
		}

		image.URL = models.ImageURL(image.Id)
		images = append(images, image)
	}

	return images, nil
}

func (repo BlogRepository) CreateBlog(blog models.Blog, contentImageIds []int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		logger.Error.Printf("error begin transaction: %v", err.Error())
		return err
	}
	defer tx.Rollback()

	var blogId int
	err = tx.QueryRow(
		"INSERT INTO blogs (title, slug, content, blog_snippet, published, cover_image_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		blog.Title, blog.Slug, blog.Content, nullableString(blog.BlogSnippet), blog.Published, blog.CoverImageId,
	).Scan(&blogId)
	if err != nil {
		logger.Error.Printf("error insert blog: %#v", err.Error())
		return err
	}

	for _, imageId := range contentImageIds {
		if _, err := tx.Exec("INSERT INTO blog_images (blog_id, image_id) VALUES ($1, $2)", blogId, imageId); err != nil {
			logger.Error.Printf("error insert blog image mapping: %#v", err.Error())
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		logger.Error.Printf("error commit transaction: %v", err.Error())
		return err
	}

	logger.Info.Printf("blog created: %v", blogId)

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

func (repo BlogRepository) UpdateBlog(id int, blog models.Blog, contentImageIds []int) error {
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
		"INSERT INTO blogs (id, title, slug, content, blog_snippet, published, cover_image_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		id, blog.Title, blog.Slug, blog.Content, nullableString(blog.BlogSnippet), blog.Published, blog.CoverImageId, blog.CreatedAt, nullableTime(blog.UpdatedAt),
	); err != nil {
		logger.Error.Printf("error insert blog in update: %v", err.Error())
		return err
	}

	if _, err := tx.Exec("DELETE FROM blog_images WHERE blog_id = $1", id); err != nil {
		logger.Error.Printf("error delete blog image mappings in update: %v", err.Error())
		return err
	}

	for _, imageId := range contentImageIds {
		if _, err := tx.Exec("INSERT INTO blog_images (blog_id, image_id) VALUES ($1, $2)", id, imageId); err != nil {
			logger.Error.Printf("error insert blog image mapping in update: %#v", err.Error())
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		logger.Error.Printf("error commit transaction: %v", err.Error())
		return err
	}

	logger.Info.Printf("blog updated: %v", id)

	return nil
}
