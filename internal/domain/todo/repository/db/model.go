package db

import (
	"time"

	"github.com/ronyelkahfi/todos/internal/domain/todo/entity"
)

type Todos struct {
	ID          int64
	title       string
	description string
	duedate     time.Time
	completed   int
	createdAt   time.Time
}

func todosFromEntity(td entity.Todo) *Todos {
	return &Todos{
		ID: td.ID,
		title: td.Title,
		description: td.Description,
		duedate: td.Duedate,
		completed: td.Completed,
		createdAt: td.CreatedAt,
	}
}

func (p *Todos) toEntity() *entity.Todo {
	return &entity.Todo{
		ID: p.ID,
		Title: p.title,
		Description: p.description,
		Duedate: p.duedate,
		Completed: p.completed,
		CreatedAt: p.createdAt,
	}
}