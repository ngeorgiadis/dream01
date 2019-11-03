package shoutcast

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// http://wiki.shoutcast.com/wiki/SHOUTcast_Radio_Directory_API

// APIClient ...
type APIClient struct {
	apiKey  string
	baseURL string
}

//StationList is a list of stations
type StationList struct {
	XMLName  xml.Name  `xml:"stationlist" json:"-"`
	TuneIn   TuneIn    `json:"tune_in"`
	Stations []Station `xml:"station" json:"stations"`
}

//
//TuneIn is part of stationlist
type TuneIn struct {
	XMLName xml.Name `xml:"tunein" json:"-"`
	Base    string   `xml:"base,attr" json:"base"`
}

// Station ...
type Station struct {
	XMLName xml.Name `xml:"station" json:"-"`

	Name  string `xml:"name,attr" json:"name"`
	MT    string `xml:"mt,attr" json:"mt"`
	ID    string `xml:"id,attr" json:"id"`
	BR    string `xml:"br,attr" json:"br"`
	Genre string `xml:"genre,attr" json:"genre"`
	Logo  string `xml:"logo,attr,omitempty" json:"logo"`
	CT    string `xml:"ct,attr" json:"ct"`
	LC    string `xml:"lc,attr" json:"lc"`
}

// New ...
func New(apiKey string) *APIClient {

	return &APIClient{
		baseURL: "http://api.shoutcast.com/legacy/",
		apiKey:  apiKey,
	}
}

// GetTop500 ...
func (c *APIClient) GetTop500() (*StationList, error) {

	apiPath := "Top500"
	url := fmt.Sprintf("%s%s?k=%s", c.baseURL, apiPath, c.apiKey)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)

	var list StationList

	err = xml.Unmarshal(b, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
