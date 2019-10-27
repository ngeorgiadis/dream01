package handlers

import (
	"dream01/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

var stor *storage.Storage

// GetRouter ...
func GetRouter(s *storage.Storage) *mux.Router {

	stor = s

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	//admin handlers

	// main handlers
	r.HandleFunc("/getrecords", getRecordsHandler)
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	r.HandleFunc("/ws", wsHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../../ui/dist"))))

	return r
}
