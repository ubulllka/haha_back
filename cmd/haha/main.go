package main

import (
	"haha/internal/app"
	"haha/internal/logger"
)

func main() {
	logg := logger.GetLogger()
	if err := app.Run(); err != nil {
		logg.Fatal(err)
	}
}
