package main

import (
	"mumbletest/mumblebot"
	"os"
)

func main() {
	mumbleServerAddress := os.Getenv("MUMBLETEST_SERVER_ADDR")
	mumbleUsername := os.Getenv("MUMBLETEST_USERNAME")
	mumblePassword := os.Getenv("MUMBLETEST_PASSWORD")

	mumbleBot := mumblebot.MumbleBot{
		ServerIP: mumbleServerAddress,
		Username: mumbleUsername,
		Password: mumblePassword,
	}

	if err := mumbleBot.Start(); err != nil {
		panic(err)
	}
}
