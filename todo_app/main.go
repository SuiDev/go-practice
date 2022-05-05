package main

import (
	//"fmt"
    //"log"
	//"todo_app/config"
    "todo_app/app/controllers"
	//"todo_app/app/models"
)

func main() {
    /*
        fmt.Println(config.Config.Port)
        fmt.Println(config.Config.SQLDriver)
        fmt.Println(config.Config.DbName)
        fmt.Println(config.Config.LogFile)

        log.Println("test")

        fmt.Println(models.Db)
    */

    /* 
        u := &models.User{}
        u.Name = "test"
        u.Email = "test@example.com"
        u.Password = "testtest"
        fmt.Println(u)

        u.CreateUser()
    */

    /*
        u, _ := models.GetUser(1)
        fmt.Println(u)

        u.Name = "test2"
        u.Email = "test2@example.com"
        u.UpdateUser()
        u, _ = models.GetUser(1)
        fmt.Println(u)
        
        u.DeleteUser()
        u, _ = models.GetUser(1)
        fmt.Println(u)
    */

    /*
        user, _ := models.GetUser(1)

        user.CreateTodo("First ToDo")
        user.CreateTodo("Second ToDo")
        t, _ := models.GetTodo(1)
        fmt.Println(t)
    */

    /*
        todos, _ := models.GetTodos()
        for _, v := range todos {
            fmt.Println(v)
        }
    */

    /*
        t.Content = "Update Todo"
        t.UpdateTodo()

        t, _ = models.GetTodo(3)
        t.DeleteTodo()

        todos, _ := user.GetTodosByUser()
        for _, v := range todos {
            fmt.Println(v)
        }
    */

    controllers.StartMainServer()
}
