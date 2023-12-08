package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	decoder := json.NewDecoder(r.Body)
	request := web.TodoRequestFindById{}
	err := decoder.Decode(&request)
	if err != nil { panic(err) }

	todoResponse := controller.service.FindById(r.Context(), request)
	webResponse := web.WebResponse{
		Code: 200,
		Message: "succesfully get todo with id " + strconv.Itoa(todoResponse.Id_todo),
		Data: todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil { panic(err) }
}

func (controller *todoControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	request := web.TodoRequestFindById{}
	err := decoder.Decode(&request)
	if err != nil { panic(err) }


}

func (controller *todoControllerImpl) SearchOrFindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	request := web.TodoRequestFindById{}
	err := decoder.Decode(&request)
	if err != nil { panic(err) }


}