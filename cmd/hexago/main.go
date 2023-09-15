package main

import (
	"fmt"
	"hexago/internal/infrastructure/server"
)

func main() {
	fmt.Print("hexagon")

	serverInstance := server.NewFiberServer()

	err := serverInstance.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
