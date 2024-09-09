package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Learning Go</h1>")
    })

    fmt.Println("Server is starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}