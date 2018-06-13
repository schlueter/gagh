package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", githubHookHandler())
	err := http.ListenAndServe(":9999", githubHookHandler())
	log.Fatal(err)
}

func githubHookHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(writer, "URL.Path: %s\nBody: %s", request.URL.Path, string(body))
	})
}
