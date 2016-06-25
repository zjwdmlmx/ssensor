package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"gopkg.in/redis.v3"
)

var R *redis.Client

func main() {
	R = redis.NewClient(&redis.Options{
		Addr:     "localhost:10000",
		PoolSize: 5,
	})

	for {
		result, err := R.BLPop(time.Minute, "FileCleaner").Result()
		if err == redis.Nil {
			continue
		}

		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 30)
			continue
		}

		fileInfo := strings.Split(result[1], " ")
		if len(fileInfo) != 2 {
			log.Println("recived unexcept command! [Error Arguments Number]")
			continue
		}

		file := fileInfo[0]
		var timeout int64
		if timeout, err = strconv.ParseInt(fileInfo[1], 10, 64); err != nil {
			log.Println("recived unexcept command! [Error Arguments]")
		}

		rm := func() {
			log.Println("Remove file:", file)
			if err := exec.Command("rm", file).Run(); err != nil {
				log.Println(err)
			}
		}

		if timeout > 0 {
			time.AfterFunc(time.Duration(timeout)*time.Second, rm)
		} else {
			rm()
		}
	}
}
