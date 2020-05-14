package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("==== Starting web server ====")
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	customName := os.Getenv("CUSTOM_NAME")
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "How are you ", "Mr. ", customName, "?")
	return
}
