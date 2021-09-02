package main

import (
	"crypto/tls"
	"mumbletest/mumblebot"
	"os"
)

func main() {
	mumbleServerAddress := os.Getenv("MUMBLETEST_SERVER_ADDR")
	mumbleUsername := os.Getenv("MUMBLETEST_USERNAME")
	mumblePassword := os.Getenv("MUMBLETEST_PASSWORD")

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	mumbleBot := mumblebot.MumbleBot{
		ServerIP:  mumbleServerAddress,
		Username:  mumbleUsername,
		Password:  mumblePassword,
		TLSConfig: tlsConfig,
	}

	if err := mumbleBot.Start(); err != nil {
		panic(err)
	}
}
