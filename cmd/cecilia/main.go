package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/test", TestHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))

	fmt.Print("asd")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test-handler"))
}
