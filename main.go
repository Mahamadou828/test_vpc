package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

func home(w http.ResponseWriter, r *http.Request) {
	url := "https://google.com"
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("can't fetch google.com: %v", err))
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("this is a simple web server %s, %v", r.URL.Path, string(b)))
}

func main() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}

}
