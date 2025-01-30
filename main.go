package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets", http.FileServer(http.Dir("/assets")))
	server := http.Server{Addr: ":8080", Handler: mux}

	server.ListenAndServe()
}

// srv := &http.Server{
// 	Addr:    ":" + port,
// 	Handler: mux,
// }

// log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
// log.Fatal(srv.ListenAndServe())
