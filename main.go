package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome!")
}

func TestIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test Index!")
}

func TestShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testId := vars["testId"]
	fmt.Fprintln(w, "Test show:", testId)
}
