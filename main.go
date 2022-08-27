package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/miyazi777/go-clean-arch/driver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
