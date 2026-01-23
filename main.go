package main

import "github.com/adm87/msgen/cmd"

var version = "0.0.0-unreleased"

func main() {
	if err := cmd.MSGen(version).Execute(); err != nil {
		panic(err)
	}
}
