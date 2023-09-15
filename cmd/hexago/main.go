package main

import (
	"fmt"
)

func main() {
	fmt.Print("hexago")

	serverInstance := initialize()

	err := serverInstance.Start()
	if err != nil {
		panic(err)
	}
}
