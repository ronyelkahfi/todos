package service

import (
	"context"

	todoRepo "github.com/ronyelkahfi/todos/internal/domain/todo/repository"
)

type Service interface {
	Create(ctx context.Context, todo CreateTodoRequest) (*Todo, int, error)
	GetAll(ctx context.Context) (*[]Todo, int, error)
	GetByID(ctx context.Context, id int64) (*Todo, int, error)
	UpdateStatusCompleteByID(ctx context.Context, id int64) (*Todo,int, error)
}

type service struct{
	todo todoRepo.Repository
}

// New to create new service.
func New(	todo todoRepo.Repository) *service {
	return &service{
		todo: todo,
	}
}