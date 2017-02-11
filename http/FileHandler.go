package http

import (
	"net/http"
	"log"
)

type FileHandler struct {
	http.Handler
}

func (h FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for %s", r.RequestURI)
	h.Handler.ServeHTTP(w, r)
}

func NewFileHandler(path string) FileHandler {
	return FileHandler{Handler: http.FileServer(http.Dir(path))}
}