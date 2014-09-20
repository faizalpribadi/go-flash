package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/setFlashMessage", set)
	http.HandleFunc("/getFlashMessage", get)
	fmt.Println("Server Running ...... ")
	http.ListenAndServe("localhost:2222", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	fm := []byte("ini flash message")
	setFlashMessage(w, "message", fm)
}

func get(w http.ResponseWriter, r *http.Request) {
	fm, err := getFlashMessage(w, r, "message")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fm == nil {
		fmt.Fprint(w, "tidak ada flash message")
		return
	}

	fmt.Fprintf(w, "%s", fm)
}
