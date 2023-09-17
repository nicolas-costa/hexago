package main

import (
	"fmt"
)

// @title Hexago API
// @version 1.0
// @contact.name API Support
// @license.name Apache 2.0
// @host localhost:3000
// @BasePath /
func main() {
	fmt.Print("hexago")

	serverInstance := initialize()

	err := serverInstance.Start()
	if err != nil {
		panic(err)
	}
}
