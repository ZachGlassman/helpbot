package main

import (
	"context"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type BaseData struct {
	Text string
}

func pullRequestHandler(w http.ResponseWriter, r *http.Request, client *github.Client, ctx *context.Context) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p github.PullRequestEvent
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	// now handle p
	switch act := *p.Action; act {
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

// Look for /merge in text
func commentMerge(s string) bool {
	var mergeRegExp = regexp.MustCompile(`\/merge`)
	matched := mergeRegExp.MatchString(s)
	return matched
}

// Get type of comment author
type CommentAuthor struct {
	Comment struct {
		AuthorAssociation string `json:"author_association"`
	} `json:"comment"`
}

func pullCommentHandler(w http.ResponseWriter, r *http.Request, client *github.Client, ctx *context.Context) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p github.IssueCommentEvent
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	// now handle p
	act := *p.Action
	if act != "deleted" {
		merg := commentMerge(*p.Comment.Body)
		// only merge if request from proper person
		if merg {
			var auth CommentAuthor
			err = json.Unmarshal(body, &auth)
			if err != nil {
				log.Println(err)
				panic(err)
			}
			switch auth.Comment.AuthorAssociation {
			case
				"OWNER",
				"COLLABORATOR",
				"MEMBER":
				tail := strings.TrimPrefix(*p.Issue.URL, "https://api.github.com/repos/")
				owner := strings.Split(tail, "/")[0]
				repo := strings.Split(tail, "/")[1]
				number, _ := strconv.ParseInt(strings.Split(tail, "/")[3], 10, 64)
				if err != nil {
					log.Println(err)
				}
				message := "Merging away, authorized by " + *p.Comment.User.Login
				client.PullRequests.Merge(*ctx, owner, repo, int(number), message, nil)
			}
		}
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
	http.HandleFunc("/issue_comment", func(w http.ResponseWriter, r *http.Request) {
		pullCommentHandler(w, r, client, &ctx)
	})
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
