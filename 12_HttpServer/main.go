package main

import "httpServer/pkg/server"

func main() {
	serv := server.NewServer()
	serv.Run()
}
