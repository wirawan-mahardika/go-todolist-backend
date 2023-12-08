package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"todolist/exception"
	"todolist/helper"
	"todolist/model/entity"
)

type TodoRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Todo, error)
	Create(ctx context.Context, tx *sql.Tx, todo entity.Todo) (entity.Todo, error)
	SearchOrFindAll(ctx context.Context, tx *sql.Tx, activity string) ([]entity.Todo, error)
	Update(ctx context.Context, tx *sql.Tx, todo entity.Todo) (entity.Todo, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}

type todoRepositoryImpl struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{}
}

func (repo *todoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Todo, error) {
	script := "SELECT id, activity, finish_target, created_at FROM todo WHERE id = ? LIMIT 1;"
	row, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)

	todo := entity.Todo{}
	if row.Next() {
		err = row.Scan(&todo.Id_todo, &todo.Activity, &todo.Finish_target, &todo.Created_at)
		helper.PanicIfError(err)
	} else {
		panic(exception.NewNotFoundError(errors.New("todo tidak ditemukan")))
	}
	return todo, nil
}

func (repo *todoRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, todo entity.Todo) (entity.Todo, error) {
	script := "INSERT INTO todo(activity, finish_target) VALUES(?,?);"
	result, err := tx.ExecContext(ctx, script, todo.Activity, todo.Finish_target)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todo, err = repo.FindById(ctx, tx, int(id))
	helper.PanicIfError(err)

	return todo, nil
}

func (repo *todoRepositoryImpl) SearchOrFindAll(ctx context.Context, tx *sql.Tx, activity string) ([]entity.Todo, error) {
	var rows *sql.Rows
	var err error
	if activity != "" {
		script := `SELECT id, activity, finish_target, created_at FROM todo WHERE activity LIKE ?;`
		searchTerm := "%" + activity + "%"
		rows, err = tx.QueryContext(ctx, script, searchTerm)
	} else {
		script := "SELECT id, activity, finish_target, created_at FROM todo;"
		rows, err = tx.QueryContext(ctx, script)
	}

	helper.PanicIfError(err)

	todos := []entity.Todo{}
	for rows.Next() {
		todo := entity.Todo{}
		err = rows.Scan(&todo.Id_todo, &todo.Activity, &todo.Finish_target, &todo.Created_at)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	if len(todos) < 1 {
		panic(exception.NewNotFoundError(errors.New("todo tidak ditemukan")))
	}

	return todos, nil
}

func (repo *todoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo entity.Todo) (entity.Todo, error) {
	script := "UPDATE todo SET "
	args := []any{}
	if todo.Activity != "" {
		script += "activity = ?"
		args = append(args, todo.Activity)
	}
	defaultTime := time.Time{}
	if todo.Finish_target != defaultTime {
		if len(args) > 0 {
			script += ", "
		}
		script += "finish_target = ?"
		args = append(args, todo.Finish_target)
	}

	script += " WHERE id = ?;"
	args = append(args, todo.Id_todo)

	_, err := tx.ExecContext(ctx, script, args...)
	helper.PanicIfError(err)

	todo, err = repo.FindById(ctx, tx, todo.Id_todo)
	helper.PanicIfError(err)

	return todo, nil
}

func (repo *todoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	script := "DELETE FROM todo WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	helper.PanicIfError(err)

	return nil
}
