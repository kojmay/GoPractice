package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/kojmay/GoPractice/2.pprof/cpu/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/kojmay"))
		}
	}()

	http.ListenAndServe(":6060", nil)
}
