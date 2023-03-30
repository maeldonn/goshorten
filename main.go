package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/maeldonn/url-shortener/api"
	"github.com/maeldonn/url-shortener/middlewares"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "The port used by the server")
	flag.Parse()
}

func main() {
	listenAddr := fmt.Sprintf(":%d", port)

	http.Handle("/", middlewares.Logger(api.NewURLHandler()))

	log.Printf("Server started on port %s\n", listenAddr)
	http.ListenAndServe(listenAddr, nil)
}
