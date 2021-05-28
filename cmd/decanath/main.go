package main

import (
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/handlers"
)

func main() {
	fmt.Println("Starting server...")
	handlers.Handle()
}
