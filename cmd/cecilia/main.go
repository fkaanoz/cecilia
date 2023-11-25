package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var release string

func main() {
	fmt.Print("release is set to : ", release)

	http.HandleFunc("/test-25", TestHandler)
	http.HandleFunc("/healthy", HealthHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test-handler"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("health check is done at : ", time.Now().UTC())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok!"))
}
