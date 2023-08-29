package repository

type User interface {
}

type Segment interface {
}

type Repository struct {
	User
	Segment
}

func NewRepository() *Repository {
	return &Repository{}
}
