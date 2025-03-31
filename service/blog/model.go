package blog

type BlogModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BlogRepositoryInterface interface {
	GetAllBlogs() (*[]BlogModel, error)
	CreateBlog(*BlogModel) (*BlogModel, error)
	GetOneBlog(id string) (*BlogModel, error)
}

type BlogQueryInterface interface {
	readAll()
}
