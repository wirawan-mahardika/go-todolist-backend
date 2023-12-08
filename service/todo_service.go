package service

import (
	"context"
	"database/sql"
	"fmt"
	"todolist/helper"
	"todolist/model/entity"
	"todolist/model/web"
	"todolist/repository"
)

type TodoService interface {
	FindById(ctx context.Context, request web.TodoRequestFindById) web.TodoResponse
	Create(ctx context.Context, request web.TodoRequestCreate) web.TodoResponse
	SearchOrFindAll(ctx context.Context, activity string) []web.TodoResponse
}

func NewTodoService(repo repository.TodoRepository, db *sql.DB) TodoService {
	return &todoServiceImpl{Repo: repo, DB: db}
}

type todoServiceImpl struct {
	Repo repository.TodoRepository
	DB   *sql.DB
}

func (service *todoServiceImpl) FindById(ctx context.Context, request web.TodoRequestFindById) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.Repo.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	webResponse := web.TodoResponse{
		Id_todo:      todo.Id_todo,
		Activity:     todo.Activity,
		FinishTarget: todo.Finish_target,
		CreatedAt:    todo.Created_at,
	}

	return webResponse
}

func (service *todoServiceImpl) Create(ctx context.Context, request web.TodoRequestCreate) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	fmt.Println(request)

	todo := entity.Todo{
		Activity:      request.Activity,
		Finish_target: request.FinishTarget,
	}

	todo, err = service.Repo.Create(ctx, tx, todo)
	helper.PanicIfError(err)

	webResponse := web.TodoResponse{
		Id_todo:      todo.Id_todo,
		Activity:     todo.Activity,
		FinishTarget: todo.Finish_target,
		CreatedAt:    todo.Created_at,
	}

	return webResponse
}

func (service *todoServiceImpl) SearchOrFindAll(ctx context.Context, activity string) []web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todos, err := service.Repo.SearchOrFindAll(ctx, tx, activity)
	helper.PanicIfError(err)

	webResponse := make([]web.TodoResponse, 0, len(todos))
	for _, todo := range todos {
		webResponse = append(webResponse, web.TodoResponse{
			Id_todo:      todo.Id_todo,
			Activity:     todo.Activity,
			FinishTarget: todo.Finish_target,
			CreatedAt:    todo.Created_at,
		})
	}

	return webResponse
}
