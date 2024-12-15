package handlers

import (
	"Todolist/internal/db/models"
	"database/sql"
	"fmt"
)

const (
	ErrorAddTask = "Error inserting task into database: %v"
)

func AddTask(DB *sql.DB, task *models.Tasks) error {
	query := `INSERT INTO Tasks(nametask,discriprion) VALUES (1$,2$)`

	_, err := DB.Exec(query, task.Nametask, task.Discription)
	if err != nil {
		return fmt.Errorf(ErrorAddTask, err)
	}

	return nil
}
