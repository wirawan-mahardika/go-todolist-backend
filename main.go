package main

import (
	"net/http"
	"todolist/app"
	"todolist/controller"
	"todolist/exception"
	"todolist/repository"
	"todolist/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.GetDBConnection()
	defer db.Close()
	router := httprouter.New()

	todoRepo := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepo, db)
	todoController := controller.NewTodoController(todoService)

	router.GET("/api/v1/todo", todoController.SearchOrFindAll)
	router.GET("/api/v1/todo/:id_todo", todoController.FindById)
	router.POST("/api/v1/todo", todoController.Create)
	router.PATCH("/api/v1/todo/:id_todo", todoController.Update)
	router.DELETE("/api/v1/todo/:id_todo", todoController.Delete)
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:1000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
