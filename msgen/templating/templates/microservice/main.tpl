{{- template "disclaimer.go.edit" }}
package main

import (
    "flag"

    "{{ .MSGenConfig.ModuleUrl }}/server"
    "{{ .MSGenConfig.ModuleUrl }}/server/generated"
)

type CmdArgs struct {
    Port int
}

func main() {
    args := parseArgs()

    s := generated.NewServer(server.New{{ .Spec.Name }}Controller())
    s.SetPort(args.Port)

    if err := s.Start(); err != nil {
        panic(err)
    }
}

func parseArgs() *CmdArgs {
    args := &CmdArgs{
        Port: 8080,
    }

    flag.IntVar(&args.Port, "port", args.Port, "Port to listen on")
    flag.Parse()

    return args
}