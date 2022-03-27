package main

import (
	"github.com/svistopljaska/blogtask/api"
	"net/http"
	"time"
)

func main() {
	r := api.InitRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
