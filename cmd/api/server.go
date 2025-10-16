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
		fmt.Println(r.Method)

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello get method from teachers route"))
			fmt.Println("Hello method get on teachers routes")
		case http.MethodPatch:
			w.Write([]byte("Hello patch method from teachers route"))
			fmt.Println("Hello method patch on teachers routes")
		case http.MethodPost:
			w.Write([]byte("Hello post method from teachers route"))
			fmt.Println("Hello method post on teachers routes")
		case http.MethodPut:
			w.Write([]byte("Hello put method from teachers route"))
			fmt.Println("Hello method put on teachers routes")
		case http.MethodDelete:
			w.Write([]byte("Hello delete method from teachers route"))
			fmt.Println("Hello method delete on teachers routes")
		default:
			w.Write([]byte("Hello from nothing method haahah"))
		}
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
