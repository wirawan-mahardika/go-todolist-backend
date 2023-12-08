package exception

import (
	"encoding/json"
	"log"
	"net/http"
	"todolist/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if notFound(w, r, i) {
		return
	}

	internalServerError(w, r, i)
}

func notFound(w http.ResponseWriter, r *http.Request, i interface{}) bool {
	exception, ok := i.(notFoundException)
	if ok {
		response := web.WebResponse{
			Code:    exception.Code,
			Message: exception.Message,
			Data:    exception.Data,
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(response)
		if err != nil {
			panic(err)
		}

		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, i interface{}) {
	log.Println(i)

	response := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data:    nil,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}
