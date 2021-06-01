package main

// Example requests: test.http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Future3")
}

func data(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"message": "Hello",
		"time": time.Now().Format(time.RFC822),
	}

	serializedData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(serializedData)

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/data", data)

	http.ListenAndServe(":8080", nil)
}