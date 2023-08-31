package service

import (
	"github.com/chigatu/service/entity"
	"github.com/chigatu/service/pkg/repository"
)

type Segment interface {
	Create(slug string) (uint, error)
	Delete(slug string) error
}

type User interface {
	ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*entity.ManageUserToSegmentsResponse, error)
	GetUserSegments(userId uint) ([]string, error)
}

type Service struct {
	Segment
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Segment: NewSegmentService(repo.Segment),
		User:    NewUsersSegmentsService(repo.User),
	}
}
