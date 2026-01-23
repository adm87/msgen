package main

import (
	"os"

	"github.com/adm87/msgen/cmd"
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
	println(msg)
	os.Exit(code)
}
