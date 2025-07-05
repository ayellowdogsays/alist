package main

import (
  "net/http"
  "os"
  "alist/server"
)

func main() {
  server.Init()
  http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
