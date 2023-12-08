package web

import "time"

type TodoRequestFindById struct {
	Id int
}

type TodoRequestCreate struct {
	Activity     string
	FinishTarget time.Time
}

type TodoRequestSearchOrFindAll struct {
	Activity     string
}

type TodoResponse struct {
	Id_todo      int		`json:"id_todo"`
	Activity     string 	`json:"activity"`
	FinishTarget time.Time		`json:"finish_target"`
	CreatedAt    time.Time  `json:"created_at"`
}