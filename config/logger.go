package config

import (
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/julienschmidt/httprouter"
)

//LoggerMiddleware API
func LoggerMiddleware(next httprouter.Handle) httprouter.Handle {
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

//Logger API
func Logger(res http.ResponseWriter, request *http.Request, p httprouter.Params) error {
	start := time.Now()

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	color.Set(color.FgCyan)

	log.Printf("%s\t%s\t%s", yellow(request.Method), yellow(request.RequestURI), green(time.Since(start)))

	color.Unset()

	return nil
}
