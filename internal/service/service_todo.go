package service

import (
	"context"
	"net/http"
	"time"

	"github.com/ronyelkahfi/todos/internal/domain/todo/entity"
	"github.com/ronyelkahfi/todos/internal/errors"
	"github.com/ronyelkahfi/todos/internal/utils"
)

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueDate string `json:"duedate" validate:"required"`
}
type Todo struct {
	ID int64 `json:id`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate time.Time `json:"duedate"`
	Completed int `json_completed`
}
func (s *service) Create(ctx context.Context, data CreateTodoRequest) (*Todo, int, error){
	if err := utils.Validate(&data); err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(ctx, err)
	}
	var dataTodo entity.Todo
	dataTodo.Title = data.Title
	dataTodo.Description = data.Description
	layout := "2006-01-02" // This layout represents "YYYY-MM-DD".

	// Parse the string into a time.Time value using the specified layout.
	duedate, err := time.Parse(layout, data.DueDate)
	if err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(ctx, err)
	}
	dataTodo.Duedate = duedate

	todo, _, err := s.todo.Insert(ctx, dataTodo)
	if err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(ctx, err)
	}
	return &Todo{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		DueDate: todo.Duedate,
		Completed: todo.Completed,
	}, http.StatusCreated, nil
}
func (s *service) GetAll(ctx context.Context) (*[]Todo, int, error){
	todos, code, err := s.todo.GetAll(ctx);
	if err != nil {
		return nil, code, errors.Wrap(ctx, err)
	}
	var todosResponse []Todo
	for _, data := range *todos{
		todosResponse = append(todosResponse, Todo{
			ID: data.ID,
			Title: data.Title,
			Description: data.Description,
			DueDate: data.Duedate,
			Completed: data.Completed,
		})
	}
	return &todosResponse, http.StatusOK, nil
}
func (s *service) GetByID(ctx context.Context, id int64) (*Todo, int, error){
	todo, code, err := s.todo.GetByID(ctx,id)
	if err != nil {
		return nil, code, errors.Wrap(ctx, err)
	}
	return &Todo{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		DueDate: todo.Duedate,
		Completed: todo.Completed,
	}, http.StatusCreated, nil
}
func (s *service) UpdateStatusCompleteByID(ctx context.Context, id int64) (*Todo, int, error){
	todo, code, err := s.todo.UpdateStatusCompleteByID(ctx,id)
	if err != nil {
		return nil, code, errors.Wrap(ctx, err)
	}
	return &Todo{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		DueDate: todo.Duedate,
		Completed: todo.Completed,
	}, http.StatusCreated, nil
}