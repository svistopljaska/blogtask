package main

import (
	"blogtask/api"
	"net/http"
	"time"
)

func main() {
	r := api.InitRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
