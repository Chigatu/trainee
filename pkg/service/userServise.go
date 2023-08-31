package service

import (
	"github.com/chigatu/service/entity"
	"github.com/chigatu/service/pkg/repository"
)

type UsersSegmentsService struct {
	repo repository.User
}

func NewUsersSegmentsService(repo repository.User) *UsersSegmentsService {
	return &UsersSegmentsService{repo: repo}
}

func (service *UsersSegmentsService) ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*entity.ManageUserToSegmentsResponse, error) {
	return service.repo.ManageUserToSegments(slugsToAdd, slugsToRemove, userId)
}
func (service *UsersSegmentsService) GetUserSegments(userId uint) ([]string, error) {
	return service.repo.GetUserSegments(userId)
}
