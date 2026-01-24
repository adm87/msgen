package main

import (
	"os"

	"github.com/adm87/msgen/cmd"
	"github.com/adm87/msgen/log"
)

var version = "0.0.0-unreleased"

func main() {
	c, err := cmd.MSGen(version)

	if err != nil {
		exit(1, err.Error())
	}

	if err := c.Execute(); err != nil {
		exit(1, err.Error())
	}
}

func exit(code int, msg string) {
	log.Error(msg)
	os.Exit(code)
}
