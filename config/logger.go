package config

import (
  "log"
  "net/http"
  "time"
  "github.com/julienschmidt/httprouter"
  "github.com/fatih/color"
)

func Logger(next httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    start := time.Now()

    yellow := color.New(color.FgYellow).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()

    color.Set(color.FgCyan)

    log.Printf("%s\t%s\t%s", yellow(r.Method), yellow(r.RequestURI), green(time.Since(start)))

    color.Unset()

    next(w, r, p)
  }
}
