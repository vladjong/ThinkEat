package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/vladjong/ThinkEat/internal/item"
)

func main() {
	router := httprouter.New()
	handler := item.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.Serve(listener))
}
