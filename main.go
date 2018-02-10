package main

import (
	"context"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type PullRequestEvent struct {
	Action       string              `json:"action,omitempty"`
	Number       int                 `json:"number,omitempty"`
	PullRequest  *github.PullRequest `json:"pull_request,omitempty"`
	Repository   *github.Repository  `json:"repository,omitempty"`
	Sender       *github.User        `json:"sender,omitempty"`
	Installation struct {
		ID string `json:"id,omitempty"`
	} `json:"installation,omitempty"`
}

type BaseData struct {
	Text string
}

func pullRequestHandler(w http.ResponseWriter, r *http.Request, client *github.Client, ctx *context.Context) {
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
		comment := &github.IssueComment{
			Body: github.String("Hi, I am helpbot, I will be managing this PR!"),
		}
		tail := strings.TrimPrefix(*p.PullRequest.URL, "https://api.github.com/repos/")
		owner := strings.Split(tail, "/")[0]
		repo := strings.Split(tail, "/")[1]
		number, _ := strconv.ParseInt(strings.Split(tail, "/")[3], 10, 64)
		_, resp, err := client.Issues.CreateComment(*ctx, owner, repo, int(number), comment)
		if err != nil {
			log.Println(owner, repo, number)
			log.Println(resp)
			log.Println(err)
		}

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

func baseHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))
	t := BaseData{Text: "Hello, this is just a webpage"}
	tmpl.Execute(w, t)
}

func main() {
	var token = os.Getenv("TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/pull_request", func(w http.ResponseWriter, r *http.Request) {
		pullRequestHandler(w, r, client, &ctx)
	})
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
