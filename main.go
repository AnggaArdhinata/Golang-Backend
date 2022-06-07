package main

import (
	"backend-golang-try/src/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env error")
	}

	fmt.Println("Vehicle Rental Service Running @PORT: 8080")
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":8080", mainRoute); err != nil {
		log.Fatal("Aplikasi gagal dijalankan")
	}

}
