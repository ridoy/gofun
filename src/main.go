package main

import (
  "fmt"
  "log"
  "net/http"
)


func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Success")
  })

  port := ":3000"
  log.Printf("listening on port %s", port)
  log.Fatal(http.ListenAndServe(port, nil))
}

