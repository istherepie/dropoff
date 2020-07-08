package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	j := &JSONResponse{200, "This is index"}

	j.Response(w)

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ALLOW CORS")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(15000000)

	uploadFile, handler, err := r.FormFile("uploadfile")
	defer uploadFile.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond to client
	j := &JSONResponse{200, "Upload completed"}
	j.Response(w)

	// Store file locally
	// TODO: Store in the /tmp directory, however this should be configurable
	fileName := fmt.Sprintf("/tmp/%s", handler.Filename)

	localFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer localFile.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Copy uploaded file
	io.Copy(localFile, uploadFile)

}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	return mux
}
