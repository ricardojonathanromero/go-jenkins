package main

import "github.com/ricardojonathanromero/go-jenkins/server"

func main() {
	// configure paths and middlewares
	srv := server.NewRouter()

	// start service to listen on port 8090
	srv.Logger.Fatal(srv.Start(":8090"))
}
