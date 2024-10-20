package main

import (
	"net/http"
	"time"

	"github.com/fadewalk/go-programming-tour-book/blog-service/internal/routers"
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
