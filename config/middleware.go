package config

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Middlewares type
type Middlewares func(res http.ResponseWriter, request *http.Request, p httprouter.Params) error

//Middleware golang
func Middleware(m ...Middlewares) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for _, middleware := range m {
			err := middleware(w, r, p)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}
