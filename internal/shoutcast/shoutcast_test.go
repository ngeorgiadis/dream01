package shoutcast

import (
	"dream01/internal/config"
	"fmt"
	"os"
	"testing"
)

func TestShoutcastTop500(t *testing.T) {

	dir, _ := os.Getwd()
	fmt.Println(dir)

	c, _ := config.New("../../cmd/server/settings.cfg")

	client := New(c.APIKey)

	stations, err := client.GetTop500()
	if err != nil {
		t.FailNow()
	}

	for _, s := range stations.Stations {
		fmt.Println(s.Name)
	}

}
