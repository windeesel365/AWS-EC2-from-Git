package main

import (
	"fmt"
	"log"
	"multgo/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
