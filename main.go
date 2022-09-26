package main

import (
	"curdapis/models"
	"curdapis/router"
	"fmt"
)

func main() {
	fmt.Print("starting the api server\n")
	models.Connect("root:admin@123@tcp(127.0.0.1:3306)/students")
	models.Migrate()
	router.ApiRouters()
	fmt.Print("Application completed...")

}
