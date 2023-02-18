package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesfull\n")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Adress = %s\n", adress)
}
func hellohandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "Get" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}
func main() {
	fileserver := http.FileServer(http.Dir("./src"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandle)
	http.HandleFunc("/hello", hellohandle)
	fmt.Println("starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
