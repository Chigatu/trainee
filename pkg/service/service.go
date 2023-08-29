package service

import "github.com/chigatu/service/pkg/repository"

type User interface {
}

type Segment interface {
}

type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
