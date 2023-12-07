package entity

import "time"

type Todo struct {
	Id_todo       int    `json:"id_todo"`
	Activity      string `json:"activity"`
	Finish_target time.Time `json:"finish_target"`
	Created_at    time.Time `json:"created_at"`
}