package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("==== Starting web server ====")
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "so far, so good")
	return
}
