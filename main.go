package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/whatDeepak/ShellSage/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
