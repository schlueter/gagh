package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const port = ":9999"

func main() {
	router := http.NewServeMux()
	router.Handle("/", githubHookHandler())
	err := http.ListenAndServe(port, githubHookHandler())
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

type GitHubEvent struct {
	Repository Repository `json:"repository"`
	// Sender Sender `json:"sender"`
}

type Repository struct {
	Name      string `json:"name"`
	Full_name string `json:"full_name"`
}

type PullRequestEvent struct {
	GitHubEvent
	Action       string      `json:"action"`
	Number       int         `json:"number"`
	Pull_request PullRequest `json:"pull_request"`
	Label        string      `json:"label"`
}

type PullRequest struct {
	Html_url string `json:"html_url"`
	Head     GitRef `json:"head"`
}

type GitRef struct {
	Label string `json:"label"`
}

func handle_webhook(body []byte) {
	var pull_request_event PullRequestEvent
	err := json.Unmarshal(body, &pull_request_event)
	if err != nil {
		log.Print("error:", err)
		return
	}
	if pull_request_event.Action == "" {
		log.Print("Noop: Recieved event without an action to act on.")
		return
	}
	log.Print("Action: Handling %s", pull_request_event.Action)
	var pull_request_number = strconv.Itoa(pull_request_event.Number)
	switch action := pull_request_event.Action; action {
	case "labeled":
		fmt.Printf("Labeled PR #%s %s in %s\n", pull_request_number, pull_request_event.Label, pull_request_event.Repository.Name)
	case "unlabeled":
		fmt.Printf("Unlabeled PR #%s %s in %s\n", pull_request_number, pull_request_event.Label, pull_request_event.Repository.Name)
	default:
		fmt.Printf("Unknown action %s\n", action)
	}
}
