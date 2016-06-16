package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/zjwdmlmx/ssensor/router"
)

func main() {
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, router.R))
}
