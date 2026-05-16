package main

import "backend-coding-challenge/server"


func main() {
	srv := server.New("config.yml")

	// init todo module
	srv.Todo()

	// init category module
	srv.Category()

	// migrate
	srv.Migrate()

	srv.Run()

}