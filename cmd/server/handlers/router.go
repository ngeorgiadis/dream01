package handlers

import (
	"dream01/internal/shoutcast"
	"dream01/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

var stor *storage.Storage
var apiClient *shoutcast.APIClient

// GetRouter ...
func GetRouter(s *storage.Storage, sc *shoutcast.APIClient) *mux.Router {

	stor = s
	apiClient = sc

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	//admin handlers

	// main handlers
	r.HandleFunc("/top", getTop500)
	r.HandleFunc("/source/{id}", getStreamSourceByID)
	r.HandleFunc("/fav", getFavourites)

	r.HandleFunc("/getrecords", getRecordsHandler)
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	r.HandleFunc("/ws", wsHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../../ui/dist"))))

	return r
}
