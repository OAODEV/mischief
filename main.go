package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handleMischif(w http.ResponseWriter, r *http.Request) {
	delay := time.Duration(rand.Intn(5000))
	time.Sleep(delay * time.Millisecond)
	if rand.Intn(2) == 0 {
		http.Error(w, "Mischeif!!!", 500)
		return
	}
	fmt.Fprintf(w, "Hi I'm HERE!!")
}

func main() {
	http.HandleFunc("/", handleMischif)
	http.ListenAndServe(":8000", nil)
}
