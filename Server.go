package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/zjwdmlmx/ssensor/router"
)

func main() {
	if cpu := runtime.NumCPU(); cpu > 1 {
		runtime.GOMAXPROCS(cpu - 1)
	}

	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, router.R))
}
