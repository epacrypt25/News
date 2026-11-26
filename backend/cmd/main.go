package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "Hello Dunia")
	})

	addr := ":8080"
	fmt.Println("Server berjalan di", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
