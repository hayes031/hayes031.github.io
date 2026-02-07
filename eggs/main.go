package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/contact.html")
	})
	http.HandleFunc("/servants", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/servants.html")
	})
	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/services.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/main.html")
	})
	http.HandleFunc("/img-r", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pics/r.jpg")
	})
	http.HandleFunc("/img-damir", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pics/damir.jpg")
	})
	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pics/semen.jpg")
	})
	http.HandleFunc("/img-ivan", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pics/ivan.jpg")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
