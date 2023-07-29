package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartHttpServer(
	ipAddress string,
	port int,
	path string,
) {
	http.Handle(
		"/",
		http.StripPrefix(
			"/",
			http.FileServer(
				http.Dir(path),
			),
		),
	)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", ipAddress, port),
			nil,
		),
	)
}

func main() {
	defaultPath, _ := os.Getwd()

	var (
		ipAddress = flag.String("ip", "localhost", "ip address of the server")
		port      = flag.Int("port", 8080, "port number of the server")
		dirPath   = flag.String("dir", defaultPath, "directory path in the server machine to publish")
	)
	flag.Parse()

	log.Printf("path:%s\n", *dirPath)
	StartHttpServer(
		*ipAddress,
		*port,
		*dirPath,
	)
}
