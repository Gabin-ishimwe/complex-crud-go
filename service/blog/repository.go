package blog

import "database/sql"

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (b *BlogRepository) GetAllBlogs() (*[]BlogModel, error) {
	return nil, nil
}

func (b *BlogRepository) CreateBlog(blog *BlogModel) (*BlogModel, error) {
	return nil, nil
}

func (b *BlogRepository) GetOneBlog(id string) (*BlogModel, error) {
	return nil, nil
}
