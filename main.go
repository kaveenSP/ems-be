package main

import (
	"ems-be/api"
	"ems-be/dbOperations"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Connecting To Database")
	dbOperations.ConnectMongoDatabase()

	fmt.Println("Starting HTTP Server")
	e := echo.New()
	api.EndpointManager(e)
	e.Logger.Fatal(e.Start(":8000"))
}
