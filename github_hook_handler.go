package main

import (
	"encoding/json"
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
		handle_webhook(body)
	})
}

type IssuesEvent struct {
	Action       string
	Number       int
	Repository   Repository
	Pull_request PullRequest
}

type Repository struct {
	Name      string
	Full_name string
}

type PullRequest struct {
	Html_url string
	Head     GitRef
}

type GitRef struct {
	Label string
}

func handle_webhook(body []byte) {
	fmt.Printf("%s\n", string(body))

	var issues_event IssuesEvent
	err := json.Unmarshal(body, &issues_event)
	if err != nil {
		log.Print("error:", err)
	}
	fmt.Printf("%+v\n", issues_event)
}
