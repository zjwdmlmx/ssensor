package controllers

import (
	"io"
	"net/http"
	"strconv"
)

func DataDownloadHandler(writer http.ResponseWriter, request *http.Request) {
	xrange := request.Header.Get("Range")

	var offset int64 = 0
	var err error
	if xrange != "" {
		offset, err = strconv.ParseInt(xrange[6:len(xrange)-1], 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			io.WriteString(writer, err.Error())
			return
		}
	}

	var newPath string
	newPath, err = copyDB()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		return
	}

	writeFile(writer, newPath, offset)
}
