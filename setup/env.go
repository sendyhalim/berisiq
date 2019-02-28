package setup

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Loading env...")

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file. %v", err)
	}
}
