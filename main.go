package main

import (
	"os"
	"github.com/fullfillmentservice/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 { port = "3000"
	}
	server := service.NewServer()
	server.Run(":" + port)
}