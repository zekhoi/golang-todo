package services

import (
	"github.com/zekhoi/golang-todo/app/dto"
	"github.com/zekhoi/golang-todo/app/entities"
	"github.com/zekhoi/golang-todo/app/repositories"
)

type ToDoService interface {
	GetAllToDos() []entities.ToDo
	GetToDo(id int) entities.ToDo
	CreateToDo(data dto.ToDoCreateRequest) entities.ToDo
	UpdateToDo(id int, data dto.ToDoUpdateRequest) entities.ToDo
	DeleteToDo(id int) error
}

type toDoService struct {
	repository repositories.ToDoRepository
}

func NewToDoService(repository repositories.ToDoRepository) ToDoService {
	return &toDoService{repository: repository}
}

func (s *toDoService) GetAllToDos() []entities.ToDo {
	return s.repository.FindAll()
}

func (s *toDoService) GetToDo(id int) entities.ToDo {
	return s.repository.FindByID(id)
}

func (s *toDoService) CreateToDo(data dto.ToDoCreateRequest) entities.ToDo {
	toDo := entities.ToDo{
		Title: data.Title,
	}

	return s.repository.Save(toDo)
}

func (s *toDoService) UpdateToDo(id int, data dto.ToDoUpdateRequest) entities.ToDo {
	toDo := s.repository.FindByID(id)
	toDo.Title = data.Title

	return s.repository.Update(toDo)
}

func (s *toDoService) DeleteToDo(id int) error {
	return s.repository.Delete(id)
}
