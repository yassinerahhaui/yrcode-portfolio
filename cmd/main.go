package main

import (
	"fmt"
	"log"
	"net/http"

	"yrcode/router"
)

func main() {
	port := ":80"
	mainMux := http.NewServeMux()
	mainMux.Handle("/", router.WebMux())
	mainMux.Handle("/api/", router.APIMux())
	fmt.Printf("Server working on => http://localhost%s\n", port)
	err := http.ListenAndServe(port, mainMux)
	if err != nil {
		log.Fatal("Server Stoped!")
	}
}
