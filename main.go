package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `envconfig:"SERVICE_PORT" default:"3000"`
}

func main() {
	var config Config
	processErr := envconfig.Process("myapp", &config)
	if processErr != nil {
		log.Fatalln(processErr)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))
	})
	address := fmt.Sprintf(":%d", config.Port)
	log.Println(fmt.Sprintf("Started; Listening %s", address))
	http.ListenAndServe(address, r)
}
