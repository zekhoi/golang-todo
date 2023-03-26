package repositories

import (
	"github.com/zekhoi/golang-todo/app/entities"
	"gorm.io/gorm"
)

type ToDoRepository interface {
	FindAll() []entities.ToDo
	FindByID(id int) entities.ToDo
	Save(todo entities.ToDo) entities.ToDo
	Update(todo entities.ToDo) entities.ToDo
	Delete(id int) error
}

type toDoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) ToDoRepository {
	return &toDoRepository{db: db}
}

func (r *toDoRepository) FindAll() []entities.ToDo {
	var todos []entities.ToDo
	r.db.Find(&todos)

	return todos
}

func (r *toDoRepository) FindByID(id int) entities.ToDo {
	var todo entities.ToDo
	r.db.Where("id = ?", id).First(&todo)

	if todo.ID < 1 {
		return entities.ToDo{}
	}

	return todo
}

func (r *toDoRepository) Save(todo entities.ToDo) entities.ToDo {
	var newTodo entities.ToDo

	newTodo = entities.ToDo{
		Title: todo.Title,
	}
	r.db.Create(&newTodo)

	return newTodo
}

func (r *toDoRepository) Update(todo entities.ToDo) entities.ToDo {
	var updatedTodo entities.ToDo

	r.db.Model(&updatedTodo).Where("id = ?", todo.ID).Updates(entities.ToDo{
		Title: todo.Title,
	})

	return updatedTodo
}

func (r *toDoRepository) Delete(id int) error {
	var todo entities.ToDo
	r.db.Where("id = ?", id).Delete(&todo)

	return nil
}
