package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	server := http.Server{Addr: ":8080", Handler: mux}

	server.ListenAndServe()
}

// srv := &http.Server{
// 	Addr:    ":" + port,
// 	Handler: mux,
// }

// log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
// log.Fatal(srv.ListenAndServe())
