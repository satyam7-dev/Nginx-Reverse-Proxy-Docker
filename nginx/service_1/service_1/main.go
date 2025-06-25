package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go Service!")
    })
    fmt.Println("Service 1 running on port 8081...")
    http.ListenAndServe(":8081", nil)
}
