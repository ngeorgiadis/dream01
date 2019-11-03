package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"dream01/cmd/server/handlers"

	"dream01/internal/config"
	"dream01/internal/intlog"
	"dream01/internal/lumberjack"
	"dream01/internal/shoutcast"
	"dream01/internal/storage"
)

var appConfig *config.Config
var stor *storage.Storage
var sc *shoutcast.APIClient

func init() {

	appConfig, _ = config.New("settings.cfg")

	logger := &lumberjack.Logger{
		Filename:   "./logs/application.log",
		MaxSize:    50, // megabytes
		MaxBackups: 15,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	}

	log.SetFlags(0)
	log.SetOutput(intlog.NewIntLogger(logger))

	stor = storage.New(storage.Settings{
		Path: "sdata",
	})

	sc = shoutcast.New(appConfig.APIKey)
}

func main() {

	go handlers.HandleMessages()

	r := handlers.GetRouter(stor, sc)

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", appConfig.Port),
		Handler: r,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT)

	errorChan := make(chan error, 1)
	go func() {
		intlog.Infof("starting subscriber service at address: [ %v ]", server.Addr)
		errorChan <- server.ListenAndServe()
	}()

	select {
	case err := <-errorChan:
		log.Fatal(err.Error())
	case sig := <-shutdown:
		intlog.Infof("interupted by user: [ %v ] ", sig.String())
		server.Close()
	}
}

func md5FromString(s string) string {
	k := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", k)
}
