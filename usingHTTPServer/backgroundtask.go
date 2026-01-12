package usinghttpserver

import (
	"fmt"
	"net/http"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	go func() {
		fmt.Println("Background task starterd")
	}()
	w.Write([]byte("Request handler"))
}
func BackgroundTask() {
	http.HandleFunc("/", Handler1)
	fmt.Println("Server running on port:8080")
	http.ListenAndServe(":8080", nil)
}
