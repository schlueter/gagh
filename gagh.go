package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/schlueter/gagh/internal/config"
)

const port = ":9999"

type gitHubEvent struct {
	Repository struct {
		Name      string `json:"name"`
		Full_name string `json:"full_name"`
	} `json:"repository"`
	// TODO Sender Sender `json:"sender"`
}

type pullRequestEvent struct {
	gitHubEvent
	Action       string `json:"action"`
	Number       int    `json:"number"`
	Label        string `json:"label"`
	Pull_request struct {
		Html_url string `json:"html_url"`
		Head     struct {
			Label string `json:"label"`
		} `json:"head"`
	} `json:"pull_request"`
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

func handle_labeled(pull_request_event pullRequestEvent) {
	log.Printf("Labeled PR %s#%s %s\n", pull_request_event.Repository.Name, strconv.Itoa(pull_request_event.Number), pull_request_event.Label)
}

func handle_unlabeled(pull_request_event pullRequestEvent) {
	log.Printf("Unlabeled PR %s#%s %s\n", pull_request_event.Repository.Name, strconv.Itoa(pull_request_event.Number), pull_request_event.Label)
}

func handle_webhook(body []byte) {
	var pull_request_event pullRequestEvent
	err := json.Unmarshal(body, &pull_request_event)
	if err != nil {
		log.Print("error parsing webhook payload:", err)
		return
	}
	if pull_request_event.Action == "" {
		log.Print("Noop: Recieved event without an action to act on.")
		return
	}
	log.Printf("Action: Handling %s", pull_request_event.Action)
	switch action := pull_request_event.Action; action {
	case "labeled":
		handle_labeled(pull_request_event)
	case "unlabeled":
		handle_unlabeled(pull_request_event)
	default:
		log.Printf("Unhandled action %s on PR %s#%s\n", action, pull_request_event.Repository.Name, strconv.Itoa(pull_request_event.Number))
	}
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Printf("%s\n", conf.GitHubToken)
	router := http.NewServeMux()
	router.Handle("/", githubHookHandler())
	err = http.ListenAndServe(port, githubHookHandler())
	log.Fatal(err)
}
