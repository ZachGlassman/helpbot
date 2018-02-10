package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func pull_request_handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p PullRequestEvent
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	// now handle p
	switch act := p.Action; act {
	case "assigned":
		fallthrough
	case "unassigned":
		fallthrough
	case "review_requested":
		fallthrough
	case "review_requested_removed":
		fallthrough
	case "labeled":
		fallthrough
	case "unlabeled":
		fallthrough
	case "opened":
		OpenPullRequest(&p)
	case "edited":
		fallthrough
	case "closed":
		fallthrough
	case "reopened":
		fallthrough
	default:
		return

	}
}

func base_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))
	t := BaseData{Text: "Hello, this is just a webpage"}
	tmpl.Execute(w, t)
}

func main() {
	http.HandleFunc("/", base_handler)
	http.HandleFunc("/pull_request", pull_request_handler)
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
