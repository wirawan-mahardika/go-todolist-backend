package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/model/entity"
)

type TodoRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Todo, error)
	Create(ctx context.Context, tx *sql.Tx, todo entity.Todo) (entity.Todo, error)
	SearchOrFindAll(ctx context.Context, tx *sql.Tx, activity string) ([]entity.Todo, error)
}

func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{}
}

type todoRepositoryImpl struct {}

func (repo *todoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Todo, error) {
	script := "SELECT id_todo, activity, finish_target, created_at FROM todo WHERE id = ? LIMIT 1;"
	row, err := tx.QueryContext(ctx, script, id)
	if err != nil { panic(err) }	

	todo := entity.Todo{}
	if row.Next() {
		err = row.Scan(&todo.Id_todo, &todo.Activity, &todo.Finish_target, &todo.Created_at)
		if err != nil { panic(err) }
	} else {
		return todo, errors.New("todo is not found")
	}
	
	return todo, nil
}

func (repo *todoRepositoryImpl)	Create(ctx context.Context, tx *sql.Tx, todo entity.Todo) (entity.Todo, error) {
	script := "INSERT INTO todo(activity, finish_target) VALUES(?,?);"
	result, err := tx.ExecContext(ctx, script, todo.Activity, todo.Finish_target)
	if err != nil { panic(err) }

	id, err := result.LastInsertId()
	if err != nil { panic(err) }

	todo, err = repo.FindById(ctx, tx, int(id))
	if err != nil { panic(err) }

	return todo, nil
}

func (repo *todoRepositoryImpl)	SearchOrFindAll(ctx context.Context, tx *sql.Tx, activity string) ([]entity.Todo, error) {
	var rows *sql.Rows
	var err error
	if activity != "" {
		script := "SELECT id_todo, activity, finish_target, created_at FROM todo WHERE activity = ? LIMIT 1;"
		rows, err = tx.QueryContext(ctx, script, activity)
	} else {
		script := "SELECT id_todo, activity, finish_target, created_at FROM todo;"
		rows, err = tx.QueryContext(ctx, script)
	}

	if err != nil { panic(err) }

	todos := []entity.Todo{}
	for rows.Next() {
		todo := entity.Todo{}
		err = rows.Scan(&todo.Id_todo, &todo.Activity, &todo.Finish_target, &todo.Created_at)
		if err != nil { panic(err) }
		todos = append(todos, todo)
	}

	return todos, nil
}