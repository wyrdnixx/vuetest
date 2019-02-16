package models

import (
	_ "database/sql"
	"fmt"
)

type Task struct {
	ID      int    `json:"id"`
	Surname string `json:"surname"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func PutTask(name string) string {

	fmt.Println("models-putTask-Name: " + name)

	return name
}
