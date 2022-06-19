package main

import (
	"go-twitter/server"
	"go-twitter/utils"
	"os"
)

func main() {
	app := server.New()

	PORT := os.Getenv("API_PORT")

	utils.InfoLog.Println("starting the server on port", PORT)

	if err := app.Listen(PORT); err != nil {
		utils.ErrorLog.Panic(err.Error())
	}
}
