package main

import (
	"fmt"
	"context"
	"github.com/munene02/go-microservice-orders-api/application"
)

func main(){
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil{
		fmt.Println("Failed to start app:", err)
	}
}