package main

import (
	"backend-golang/src/configs/command"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to Vehicle-Rental Back-End APP")
}
