package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"dream01/internal/intlog"

	"github.com/gorilla/mux"
)

// RadioStation ...
type RadioStation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

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

func getFavourites(w http.ResponseWriter, r *http.Request) {

	sl := []RadioStation{
		RadioStation{
			Name: "dream01",
			URL:  "http://dream01.gr:8000/",
		},
	}

	b, _ := json.Marshal(&sl)

	w.Write(b)

}

func getTop500(w http.ResponseWriter, r *http.Request) {

	sl, _ := apiClient.GetTop500()

	b, _ := json.Marshal(&sl)

	w.Write(b)

}

func getStreamSourceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	base := "/sbin/tunein-station.pls"

	// var url = `http://yp.shoutcast.com${payload.base}?id=${payload.id}`;
	url := fmt.Sprintf("http://yp.shoutcast.com%s?id=%v", base, vars["id"])

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))

	lines := strings.Split(string(b), "\n")

	for _, l := range lines {
		l = strings.TrimSpace(l)

		if strings.Index(l, "File1") > -1 {
			i := strings.Index(l, "=")

			stream := fmt.Sprintf("%v", l[i+1:])

			w.Write([]byte(stream))

			break
		}
	}

}
