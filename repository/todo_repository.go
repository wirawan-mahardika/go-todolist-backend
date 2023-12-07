package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/model/entity"
)

type todoRepository interface {
	FindById(ctx context.Context, id int) (entity.Todo, error)
	Create(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	SearchOrFindAll(ctx context.Context, activity string) ([]entity.Todo, error)
}

func NewTodoRepository(db *sql.DB) todoRepository {
	return &todoRepositoryImpl{DB: db}
}

type todoRepositoryImpl struct {
	DB *sql.DB
}

func (repo *todoRepositoryImpl) FindById(ctx context.Context, id int) (entity.Todo, error) {
	script := "SELECT id_todo, activity, finish_target, created_at FROM todo WHERE id = ? LIMIT 1;"
	row, err := repo.DB.QueryContext(ctx, script, id)
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

func (repo *todoRepositoryImpl)	Create(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	script := "INSERT INTO todo(activity, finish_target) VALUES(?,?);"
	result, err := repo.DB.ExecContext(ctx, script, todo.Activity, todo.Finish_target)
	if err != nil { panic(err) }

	id, err := result.LastInsertId()
	if err != nil { panic(err) }

	todo, err = repo.FindById(ctx, int(id))
	if err != nil { panic(err) }

	return todo, nil
}

func (repo *todoRepositoryImpl)	SearchOrFindAll(ctx context.Context, activity string) ([]entity.Todo, error) {
	var rows *sql.Rows
	var err error
	if activity != "" {
		script := "SELECT id_todo, activity, finish_target, created_at FROM todo WHERE activity = ? LIMIT 1;"
		rows, err = repo.DB.QueryContext(ctx, script, activity)
	} else {
		script := "SELECT id_todo, activity, finish_target, created_at FROM todo;"
		rows, err = repo.DB.QueryContext(ctx, script)
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