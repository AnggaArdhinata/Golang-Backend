package main

import (
	"backend-golang/src/configs/command"
	"log"
	"os"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
