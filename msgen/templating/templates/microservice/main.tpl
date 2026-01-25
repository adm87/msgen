{{- template "disclaimer.go.noedit" }}
package main

import (
    "{{ .MSGenConfig.ModuleUrl }}/server"
    "{{ .MSGenConfig.ModuleUrl }}/server/generated"
)

func main() {
    s := generated.NewServer(server.New{{ .Spec.Name }}Controller())

    if err := s.Start(); err != nil {
        panic(err)
    }
}