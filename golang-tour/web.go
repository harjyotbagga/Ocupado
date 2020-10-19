package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h3>About Us</h3>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Staring...")
	http.ListenAndServe(":3000", nil)
}
