package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist/helper"
	"todolist/model/web"
	"todolist/service"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	SearchOrFindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

func NewTodoController(service service.TodoService) TodoController {
	return &todoControllerImpl{service: service}
}

type todoControllerImpl struct {
	service service.TodoService
}

func (controller *todoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id_todo"))
	helper.PanicIfError(err)
	request := web.TodoRequestFindById{Id: id}

	todoResponse := controller.service.FindById(r.Context(), request)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "succesfully get todo with id " + strconv.Itoa(todoResponse.Id_todo),
		Data:    todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *todoControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	request := web.TodoRequestCreate{}
	err := decoder.Decode(&request)
	helper.PanicIfError(err)

	todoResponse := controller.service.Create(r.Context(), request)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "success create new todo",
		Data:    todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *todoControllerImpl) SearchOrFindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := web.TodoRequestSearchOrFindAll{Activity: r.URL.Query().Get("activity")}

	todoResponse := controller.service.SearchOrFindAll(r.Context(), request.Activity)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "succesfully get todos",
		Data:    todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
