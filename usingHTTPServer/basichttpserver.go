package usinghttpserver

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request")
	w.Write([]byte("Hello world!"))
}

func BasicHTTPServer() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server running on port:8080")
	http.ListenAndServe(":8080", nil)
}
