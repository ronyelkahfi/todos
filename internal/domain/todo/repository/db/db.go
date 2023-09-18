package db

import (
	"context"
	_errors "errors"
	"net/http"

	"github.com/ronyelkahfi/todos/internal/domain/todo/entity"
	"github.com/ronyelkahfi/todos/internal/errors"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}
func (db *DB) Insert(ctx context.Context, todo entity.Todo) (*entity.Todo, int, error) {
	tx := db.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, tx.Error)
	}
	defer tx.Rollback()

	p := todosFromEntity(todo)
	if err := tx.Create(p).Error; err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, err)
	}

	
	if err := tx.Commit().Error; err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, err)
	}

	return p.toEntity(), http.StatusOK, nil
}

// GetByID to get payment by id.
func (db *DB) GetByID(ctx context.Context, id int64) (*entity.Todo, int, error) {
	var p Todos
	if err := db.db.WithContext(ctx).Where("id = ?", id).First(&p).Error; err != nil {
		if _errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.Wrap(ctx, errors.ErrNotFoundPayment, err)
		}
		return nil, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, err)
	}
	return p.toEntity(), http.StatusOK, nil
}

// GetAll to get todo by id.
func (db *DB) GetAll(ctx context.Context) (*[]entity.Todo, int, error) {
	var listTodo []Todos
	if err := db.db.WithContext(ctx).Find(&listTodo).Error; err != nil {
		if _errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.Wrap(ctx, errors.ErrNotFoundPayment, err)
		}
		return nil, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, err)
	}
	var listTodoEntity []entity.Todo
	for  _, recordTodo := range listTodo {
		listTodoEntity = append(listTodoEntity, *recordTodo.toEntity())
	}
	return &listTodoEntity, http.StatusOK, nil
}

// UpdateStatusDoneByID to update payment status to done.
func (db *DB) UpdateStatusCompleteByID(ctx context.Context, id int64) (*entity.Todo, int, error) {
	tx := db.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return &entity.Todo{}, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, tx.Error)
	}
	defer tx.Rollback()

	query := `
		UPDATE todos
		SET complete = ?
		WHERE id = ?`

	res := tx.Exec(query, entity.StatusCompleted, id)
	
	if res.Error != nil {
		return &entity.Todo{}, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, res.Error)
	}

	if res.RowsAffected == 0 {
		return &entity.Todo{}, http.StatusNotFound, errors.Wrap(ctx, errors.ErrNotFoundPayment)
	}

	if err := tx.Commit().Error; err != nil {
		return &entity.Todo{}, http.StatusInternalServerError, errors.Wrap(ctx, errors.ErrInternalDB, err)
	}
	record, status, _ := db.GetByID(ctx, id)
	return &entity.Todo{
		ID: record.ID,
		Title: record.Title,
		Description: record.Description,
		Duedate: record.Duedate,
		Completed: record.Completed,
		CreatedAt: record.CreatedAt,
	} , status, nil
}
