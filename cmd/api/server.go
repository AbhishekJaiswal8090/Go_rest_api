package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello root route \n")
		w.Write([]byte("Hello from root route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from teacher route"))
		fmt.Println("Hello teacher Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from students route"))
		fmt.Println("Hello student Route")
	})

	fmt.Println("Server is running on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error statring the server ", err)
	}
}
