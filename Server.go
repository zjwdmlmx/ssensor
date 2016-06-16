package main

import (
	"net/http"
	"os"
	"router"

	"github.com/gorilla/handlers"
)

func main() {
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, router.R))
}
