package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONResponse struct {
	Status      int    `json:"status"`
	Description string `json:"desc"`
}

func (j *JSONResponse) Response(w http.ResponseWriter) {
	res, err := json.Marshal(j)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", res)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	j := &JSONResponse{200, "This is index"}

	j.Response(w)

}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	return mux
}
