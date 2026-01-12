package usinghttpserver

import (
	"fmt"
	"net/http"
)

func SendEmail(email string) {
	fmt.Println("Sending email to", email)
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	go SendEmail(email)

	w.Write([]byte("Email requset accepted"))
}

func BackgroundEmail() {
	http.HandleFunc("/", Handler2)
	fmt.Println("Server running on port:8080")
	http.ListenAndServe(":8080", nil)
}
