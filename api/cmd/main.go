package main

import (
	"bootcamp_api/api/server"
	"log"
)

func main() {
    srv, err := server.New()
    if err != nil {
        log.Fatalf("Failed to create server: %v", err)
    }

    //Mostrar los logs con el error y la capa en donde ocurrio
    log.Println("Listening on :8080...")
    log.Fatal(srv.Start(":8080"))
}
