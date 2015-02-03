package main

import (
	"log"
	"net/http"
)

func main() {
	dir := "/data"
	port := ":80"

	log.Printf("Serving static files from %s ...", dir)

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Printf("Listening on %s ...", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Quitting...")
}
