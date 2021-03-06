package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	fmt.Println("Running..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
