package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Body string
}

func OpenPullRequest(p *PullRequestEvent) {
	// Say Hello
	log.Println("Saying hello to new pull request")
	url := p.PullRequest.IssueURL + "/comments"
	m := Message{Body: "Hello, I am helpbot and I will help manage this PR."}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	http.NewRequest("POST", url, b)
}
