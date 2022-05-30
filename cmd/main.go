package main

import (
	"log"

	taskbar "main.go/TaskBar"
)

func main() {
	srv := new(taskbar.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while rinnong server:%s", err.Error())
	}
}
