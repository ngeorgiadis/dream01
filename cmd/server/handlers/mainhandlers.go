package handlers

import (
	"io/ioutil"
	"net/http"

	"dream01/internal/intlog"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("../../ui/dist/index.html")
	if err != nil {
		intlog.Error(err.Error())
	}
	w.Write(b)

}

func getRecordsHandler(w http.ResponseWriter, r *http.Request) {

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

}
