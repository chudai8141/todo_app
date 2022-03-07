package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@exmaple.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()

	user, _ := models.GetUser(1)
	user.CreateTodo("First Todo")

	// u, _ := models.GetUser(1)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)

	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

}
