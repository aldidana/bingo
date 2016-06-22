package render

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//ToJSON return json
func ToJSON(res http.ResponseWriter, request *http.Request, p httprouter.Params, value interface{}) {
	res.Header().Add("Content-Type", "application/json; charset=UTF-8")
	res.Header().Set("Content-Encoding", "gzip")
	res.WriteHeader(http.StatusOK)

	gz := gzip.NewWriter(res)
	defer gz.Close()

	if err := json.NewEncoder(gz).Encode(value); err != nil {
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}

//ToXML return xml
func ToXML(res http.ResponseWriter, request *http.Request, p httprouter.Params, value interface{}) {
	res.Header().Add("Content-Type", "application/xml; charset=UTF-8")
	res.Header().Set("Content-Encoding", "gzip")
	res.WriteHeader(http.StatusOK)

	gz := gzip.NewWriter(res)
	defer gz.Close()

	if err := xml.NewEncoder(gz).Encode(value); err != nil {
		res.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
