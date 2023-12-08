package main

import (
	"log"
	"net/http"
	"todolist/app"
	"todolist/controller"
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

	router.GET("/api/v1/todo", todoController.FindById)
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Println(i)
		http.Error(w, i.(string), http.StatusInternalServerError)
	}

	server := http.Server{
		Addr: "localhost:1000", 
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil { panic(err) }
}