package main

import (
	"github.com/fadewalk/go-programming-tour-book/blog-service/internal/routers"
	
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
}
