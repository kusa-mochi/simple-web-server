package main

import (
	"log"
	"net/http"
	"os"
)

func StartHttpServer(path string) {
	http.Handle(
		"/",
		http.StripPrefix(
			"/",
			http.FileServer(
				http.Dir(path),
			),
		),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	path, _ := os.Getwd()
	StartHttpServer(path)
}
