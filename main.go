package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	// use models.Db
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
