package controllers

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/zjwdmlmx/ssensor/global"
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

func writeFile(writer http.ResponseWriter, path string, offset int64) (ok bool) {
	ok = true
	file, err := os.Open(path)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		ok = false
		return
	}

	var fstat os.FileInfo
	fstat, err = file.Stat()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, err.Error())
		ok = false
		return
	}

	file_size := fstat.Size()

	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Accept-Ranges", "bytes")
	writer.Header().Set("Content-Length", fmt.Sprint(file_size-offset))
	writer.Header().Set("Content-Disposition", "attachment; filename=\"data.db\"")
	writer.Header().Set("Date", fstat.ModTime().Format("Mon, 2 Jan 2006 15:04:05 MST"))

	if offset > 0 {
		writer.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", offset, file_size-1, file_size))
		writer.WriteHeader(http.StatusPartialContent)
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	buffer := make([]byte, 4086)
	var readc int = 0

	file.Seek(offset, 0)

	for {
		readc, err = file.Read(buffer)

		if err == io.EOF || err == io.ErrNoProgress || err == io.ErrUnexpectedEOF {
			break
		}

		writer.Write(buffer[:readc])
	}

	return
}

func getFileMD5Sum(reader io.Reader) (sum string, err error) {
	hash := md5.New()

	buffer := make([]byte, 4086)
	var count int

	for {
		count, err = reader.Read(buffer)

		if err == io.EOF {
			err = nil
			sum = fmt.Sprintf("%x", hash.Sum(nil))
			return
		}
		if err != nil {
			return
		}

		hash.Write(buffer[:count])
	}
}

// copy the data base file to a temp file
func copyDB() (newPath string, err error) {
	var file *os.File
	if file, err = os.Open(global.DBPath); err != nil {
		return
	}

	defer file.Close()

	global.DB.Exec("BEGIN IMMEDIATE TRANSACTION")
	global.DB.Exec("select * from sensor_data where id=1")
	defer global.DB.Exec("COMMIT") // make sure the lock free

	var fileMd5 string
	if fileMd5, err = getFileMD5Sum(file); err != nil {
		return
	}
	file.Seek(0, 0)

	newPath = os.TempDir() + "/" + fileMd5
	_, err = os.Stat(newPath)
	if err != nil && !os.IsExist(err) {
		var newFile *os.File
		if newFile, err = os.Create(newPath); err != nil {
			return
		}

		defer newFile.Close()

		if err = global.R.RPush("FileCleaner", fmt.Sprintf("%s %d", newPath, 3600*6)).Err(); err != nil {
			return
		}

		if _, err = io.Copy(newFile, file); err != nil {
			return
		}

		log.Println("[COPY] there is database file copy operation")
	}

	return
}
