package main

import (
	"github.com/joho/godotenv"
	"haha/internal/app"
	"haha/internal/logger"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	logg := logger.GetLogger()
	if err := app.Run(logg); err != nil {
		logg.Fatal(err)
	}
}
