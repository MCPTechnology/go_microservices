package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ip   string = ""
	port int32  = 9090
)

var (
	addr    string       = fmt.Sprintf("%v:%v", ip, port)
	handler http.Handler = nil
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world")
	})

	log.Println("Server is running...")
	http.ListenAndServe(addr, handler)
}
