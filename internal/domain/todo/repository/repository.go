package repository

import (
	"context"

	"github.com/ronyelkahfi/todos/internal/domain/todo/entity"
)

type Repository interface {
	Insert(ctx context.Context, todo entity.Todo) (*entity.Todo, int, error)
	GetAll(ctx context.Context) (*[]entity.Todo, int, error)
	GetByID(ctx context.Context, id int64) (*entity.Todo, int, error)
	UpdateStatusCompleteByID(ctx context.Context, id int64) (*entity.Todo, int, error)
}
