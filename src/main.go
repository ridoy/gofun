package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
  switch string(ctx.Path()) {
  case "/hello":
    fmt.Fprintf(ctx, "Success")
  default:
    ctx.Error("Unsupported path", fasthttp.StatusNotFound)
  }
}

func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Success")
  })
  port := ":3000"
  fastHTTPPort := ":3001"
  log.Printf("listening on ports %s and %s", port, fastHTTPPort)
  /*go func() {
    log.Fatal(http.ListenAndServe(port, nil))
  }()*/
  log.Fatal(fasthttp.ListenAndServe(fastHTTPPort, fastHTTPHandler))
}

