package web

import "time"

type TodoRequestFindById struct {
	Id int
}

type TodoRequestCreate struct {
	Activity     string    `json:"activity"`
	FinishTarget time.Time `json:"finish_target"`
}

type TodoRequestSearchOrFindAll struct {
	Activity string `json:"activity"`
}

type TodoRequestUpdate struct {
	Id           int       `json:"id_todo"`
	Activity     string    `json:"activity"`
	FinishTarget time.Time `json:"finish_target"`
}

type TodoRequestDelete struct {
	Id int `json:"id_todo"`
}

type TodoResponse struct {
	Id_todo      int       `json:"id_todo"`
	Activity     string    `json:"activity"`
	FinishTarget time.Time `json:"finish_target"`
	CreatedAt    time.Time `json:"created_at"`
}
