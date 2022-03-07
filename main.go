package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// use models.Db
	fmt.Println(models.Db)

	// output config
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// setting u (User) params
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@exmaple.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// create user code
	// u.CreateUser()

	// // create todo for user
	// user, _ := models.GetUser(1)
	// user.CreateTodo("First Todo")

	// u, _ := models.GetUser(1)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)

	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// get todos with user
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")

	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}

}
