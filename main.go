package main

import (
	"os"

	"github.com/adm87/msgen/cmd"
	"github.com/adm87/msgen/styles/labels"
)

func main() {
	if err := cmd.MSGen(); err != nil {
		os.Stderr.WriteString(labels.Error(err) + "\n")
		os.Exit(1)
	}
}
