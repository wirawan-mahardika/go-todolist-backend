package service

import (
	"context"
	"database/sql"
	"todolist/helper"
	"todolist/model/entity"
	"todolist/model/web"
	"todolist/repository"
)

type TodoService interface {
	FindById(ctx context.Context, request web.TodoRequestFindById) web.TodoResponse
	Create(ctx context.Context, request web.TodoRequestCreate) web.TodoResponse
	SearchOrFindAll(ctx context.Context, activity string) []web.TodoResponse
	Update(ctx context.Context, request web.TodoRequestUpdate) web.TodoResponse
	Delete(ctx context.Context, request web.TodoRequestDelete)
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

func (service *todoServiceImpl) Update(ctx context.Context, request web.TodoRequestUpdate) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := entity.Todo{
		Id_todo:       request.Id,
		Activity:      request.Activity,
		Finish_target: request.FinishTarget,
	}
	todo, err = service.Repo.Update(ctx, tx, todo)
	helper.PanicIfError(err)

	webResponse := web.TodoResponse{
		Id_todo:      todo.Id_todo,
		Activity:     todo.Activity,
		FinishTarget: todo.Finish_target,
		CreatedAt:    todo.Created_at,
	}

	return webResponse
}

func (service *todoServiceImpl) Delete(ctx context.Context, request web.TodoRequestDelete) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Repo.Delete(ctx, tx, request.Id)
	helper.PanicIfError(err)

}
