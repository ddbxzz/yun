package main

import (
	"net/http"
	"yun/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "0.0.0.0:8080",
	}

	server.ListenAndServe()
}
