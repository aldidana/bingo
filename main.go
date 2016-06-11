package main

import (
  "net/http"
  "log"
  "fmt"
  "github.com/aldidana/bingo/config"
)

func main() {
  fmt.Print("Bingo Running on port 3000")
  log.Fatal(http.ListenAndServe(":3000", config.Router()))
}
