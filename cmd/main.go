package main

import (
	"Todolist/config"
	"Todolist/internal/db"
	"Todolist/internal/db/models"
	"Todolist/internal/handlers"
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	ErrorConnectDB = "Error connecting to database: %v"
	ErrorAddTask   = "Error add task to database: %v"
)

func main() {

	cfg := config.LoadConfig()

	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf(ErrorConnectDB, err)
	}
	defer database.Close()

	reader := bufio.NewReader(os.Stdin)

	TaskName, _ := reader.ReadString('\n')
	Discr, _ := reader.ReadString('\n')

	task := &models.Tasks{
		Nametask:    TaskName[:len(TaskName)-1],
		Discription: Discr[:len(Discr)-1],
	}

	err = handlers.AddTask(database, task)
	if err != nil {
		log.Fatalf(ErrorAddTask, err)
	}

	fmt.Println("All is ok!!!")

}
