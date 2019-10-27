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
	// r.HandleFunc("/admin", adminHandler)
	// r.HandleFunc("/admin/import", adminImportHandler)
	// r.HandleFunc("/admin/delete-completed", deleteCompletedHandler)

	// main handlers
	r.HandleFunc("/getrecords", getRecordsHandler)
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	// r.HandleFunc("/save-comment", saveCommentHandler).Methods("POST")
	// r.HandleFunc("/update-record", updateRecordHandler).Methods("POST")
	// r.HandleFunc("/update-records", updateRecordsHandler).Methods("POST")

	// r.HandleFunc("/delete-records", deleteRecordsHandler).Methods("POST")
	// r.HandleFunc("/undo-delete", undoDeleteHandler).Methods("POST")

	r.HandleFunc("/ws", wsHandler)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./ui/dist"))))
	//r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./front/public"))))

	return r
}
