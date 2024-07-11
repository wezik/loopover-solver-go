package main

import (
	"github.com/joho/godotenv"
	"loopover-solver/api"
	"os"
)

func main() {
	loadProfile()
	api.StartServer()
}

func loadProfile() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	profile := os.Getenv("LOOPOVER_SOLVER_PROFILE")
	if profile == "prod" {
		err = godotenv.Load("prod.env")
	} else {
		err = godotenv.Load("dev.env")
	}
	if err != nil {
		panic(err)
	}
}
