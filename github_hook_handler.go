package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(writer, "URL.Path: %s\nBody: %s", request.URL.Path, string(body))
	})

	err := http.ListenAndServe(":9999", handler)
	log.Fatal(err)
}
