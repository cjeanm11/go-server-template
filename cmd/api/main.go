package main

import (
	"server-template/internal/server"
)

func main() {
	srv := server.NewServer()
	srv.Start()
}
