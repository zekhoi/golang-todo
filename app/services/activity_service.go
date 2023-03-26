package services

import (
	"github.com/zekhoi/golang-todo/app/dto"
	"github.com/zekhoi/golang-todo/app/entities"
	"github.com/zekhoi/golang-todo/app/repositories"
)

type ActivityService interface {
	GetAllActivities() []entities.Activity
	GetActivity(id int) entities.Activity
	CreateActivity(data dto.ActivityCreateRequest) entities.Activity
	UpdateActivity(id int, data dto.ActivityUpdateRequest) entities.Activity
	DeleteActivity(id int) error
}

type activityService struct {
	repository repositories.ActivityRepository
}

func NewActivityService(repository repositories.ActivityRepository) ActivityService {
	return &activityService{repository: repository}
}

func (s *activityService) GetAllActivities() []entities.Activity {
	return s.repository.FindAll()
}

func (s *activityService) GetActivity(id int) entities.Activity {
	return s.repository.FindByID(id)
}

func (s *activityService) CreateActivity(data dto.ActivityCreateRequest) entities.Activity {
	activity := entities.Activity{
		Title: data.Title,
	}

	return s.repository.Save(activity)
}

func (s *activityService) UpdateActivity(id int, data dto.ActivityUpdateRequest) entities.Activity {
	activity := s.repository.FindByID(id)
	activity.Title = data.Title

	return s.repository.Update(activity)
}

func (s *activityService) DeleteActivity(id int) error {
	return s.repository.Delete(id)
}
