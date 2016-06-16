package controllers

import (
	"io"
	"net/http"

	"github.com/zjwdmlmx/ssensor/mime"
)

func writeJSON(writer http.ResponseWriter, v interface{}) (ok bool) {
	res, err := mime.JSONString(v)
	ok = true

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		ok = false
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	return
}

func readForm(writer http.ResponseWriter, request *http.Request) (ok bool) {
	err := request.ParseForm()
	ok = true
	if err != nil {
		writer.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(writer, err.Error())
		ok = false
		return
	}
	return
}
