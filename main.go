package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_url,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}

type Comment struct {
	URL       string `json:"url,omitempty"`
	HTMLURL   string `json:"html_url,omitempty"`
	ID        int    `json:"id,omitempty"`
	User      *User  `json:"user,omitempty"`
	Position  string `json:"position,omitempty"`
	Line      string `json:"line,omitempty"`
	Path      string `json:"path,omitempty"`
	CommitId  string `json:"commit_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Body      string `json:"body,omitempty"`
}

type Repository struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	FullName         string `json:"full_name,omitempty"`
	Owner            *User  `json:"owner,omitempty"`
	Private          bool   `json:"private,omitempty"`
	HTMLURL          string `json:"html_url,omitempty"`
	Description      string `json:"description,omitempty"`
	Fork             bool   `json:"fork,omitempty"`
	URL              string `json:"url,omitempty"`
	ForksURL         string `json:"forks_url,omitempty"`
	KeysURL          string `json:"keys_url,omitempty"`
	CollaboratorsURL string `json:"collaborators_url,omitempty"`
	TeamsURL         string `json:"teams_url,omitempty"`
	HooksURL         string `json:"hooks_url,omitempty"`
	IssueEventsURL   string `json:"issue_events_url,omitempty"`
	EventsURL        string `json:"events_url,omitempty"`
	AssigneesURL     string `json:"assignees_url,omitempty"`
	BranchesURL      string `json:"branches_url,omitempty"`
	TagsURL          string `json:"tags_url,omitempty"`
	BlobsURL         string `json:"blobs_url,omitempty"`
	GitTagsURL       string `json:"git_tags_url,omitempty"`
	GitRefsURL       string `json:"git_refs_url,omitempty"`
	TreesURL         string `json:"trees_url,omitempty"`
	StatusesURL      string `json:"statuses_url,omitempty"`
	LanguagesURL     string `json:"languages_url,omitempty"`
	StargazersURL    string `json:"stargazers_url,omitempty"`
	ContributorsURL  string `json:"contributors_url,omitempty"`
	SubscribersURL   string `json:"subscribers_url,omitempty"`
	SubscriptionURL  string `json:"subscription_url,omitempty"`
	CommitsURL       string `json:"commits_url,omitempty"`
	GitCommitsURL    string `json:"git_commits_url,omitempty"`
	CommentsURL      string `json:"comments_url,omitempty"`
	IssueCommentURL  string `json:"issue_comment_url,omitempty"`
	ContentsURL      string `json:"contents_url,omitempty"`
	CompareURL       string `json:"compare_url,omitempty"`
	MergesURL        string `json:"merges_url,omitempty"`
	ArchiveURL       string `json:"archive_url,omitempty"`
	DownloadsURL     string `json:"downloads_url,omitempty"`
	IssuesURL        string `json:"issues_url,omitempty"`
	PullsURL         string `json:"pulls_url,omitempty"`
	MilestonesURL    string `json:"milestones_url,omitempty"`
	NotificationsURL string `json:"notifications_url,omitempty"`
	LabelsURL        string `json:"labels_url,omitempty"`
	ReleasesURL      string `json:"releases_url,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	PushedAt         string `json:"pushed_at,omitempty"`
	GitURL           string `json:"git_url,omitempty"`
	SSHURL           string `json:"ssh_url,omitempty"`
	CloneURL         string `json:"clone_url,omitempty"`
	SvnURL           string `json:"svn_url,omitempty"`
	Homepage         string `json:"homepage,omitempty"`
	Size             int    `json:"size,omitempty"`
	StargazersCount  int    `json:"stargazers_count,omitempty"`
	WatchersCount    int    `json:"watchers_count,omitempty"`
	Language         string `json:"language,omitempty"`
	HasIssues        bool   `json:"has_issues,omitempty"`
	HasDownloads     bool   `json:"has_downloads,omitempty"`
	HasWiki          bool   `json:"has_wiki,omitempty"`
	HasPages         bool   `json:"has_pages,omitempty"`
	ForksCount       int    `json:"forks_count,omitempty"`
	MirrorURL        string `json:"mirror_url,omitempty"`
	OpenIssuesCount  int    `json:"open_issues_count,omitempty"`
	Forks            int    `json:"forks,omitempty"`
	OpenIssues       int    `json:"open_issues,omitempty"`
	Watchers         int    `json:"watchers,omitempty"`
	DefaultBranch    string `json:"default_branch,omitempty"`
}

type Commit struct {
	Label string      `json:"label,omitempty"`
	Ref   string      `json:"ref,omitempty"`
	SHA   string      `json:"sha,omitempty"`
	User  *User       `json:"user,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
}

type KeyValue struct {
	Key   string
	Value string
}

type Links struct {
	Self           *KeyValue `json:"self,omitempty"`
	HTML           *KeyValue `json:"html,omitempty"`
	Issue          *KeyValue `json:"issue,omitempty"`
	Comments       *KeyValue `json:"comments,omitempty"`
	ReviewComments *KeyValue `json:"review_comments,omitempty"`
	ReviewComment  *KeyValue `json:"review_comment,omitempty"`
	Commits        *KeyValue `json:"commits,omitempty"`
	Statuses       *KeyValue `json:"statuses,omitempty"`
}
type PullRequest struct {
	URL               string  `json:"url,omitempty"`
	ID                int     `json:"id,omitempty"`
	HTMLURL           string  `json:"html_url,omitempty"`
	DiffURL           string  `json:"diff_url,omitempty"`
	PatchURL          string  `json:"patch_url,omitempty"`
	IssueURL          string  `json:"issue_url,omitempty"`
	Number            int     `json:"number,omitempty"`
	State             string  `json:"state,omitempty"`
	Locked            bool    `json:"locked,omitempty"`
	Title             string  `json:"title,omitempty"`
	User              *User   `json:"user,omitempty"`
	Body              string  `json:"body,omitempty"`
	CreatedAt         string  `json:"created_at,omitempty"`
	UpdatedAt         string  `json:"updated_at,omitempty"`
	ClosedAt          string  `json:"closed_at,omitempty"`
	MergedAt          string  `json:"merged_at,omitempty"`
	MergeCommitSHA    string  `json:"merge_commit_sha,omitempty"`
	Assignee          string  `json:"assignee,omitempty"`
	Milestone         string  `json:"milestone,omitempty"`
	CommitsURL        string  `json:"commits_url,omitempty"`
	ReviewCommentsURL string  `json:"review_comments_url,omitempty"`
	ReviewCommentURL  string  `json:"review_comment_url,omitempty"`
	CommentsURL       string  `json:"comments_url,omitempty"`
	StatusesURL       string  `json:"statuses_url,omitempty"`
	Head              *Commit `json:"head,omitempty"`
	Base              *Commit `json:"base,omitempty"`
	Links             *Links  `json:"links,omitempty"`
	Merged            bool    `json:"merged,omitempty"`
	Mergable          bool    `json:"mergable,omitempty"`
	MergableState     string  `json:"mergable_state,omitempty"`
	MergedBy          string  `json:"merged_by,omitempty"`
	Comments          int     `json:"comments,omitempty"`
	ReviewComments    int     `json:"review_comments,omitempty"`
	Commits           int     `json:"commits,omitempty"`
	Additions         int     `json:"additions,omitempty"`
	Deletions         int     `json:"deletions,omitempty"`
	ChangedFiles      int     `json:"changed_files,omitempty"`
}

type PullRequestEvent struct {
	Action       string       `json:"action,omitempty"`
	Number       int          `json:"number,omitempty"`
	PullRequest  *PullRequest `json:"pull_request,omitempty"`
	Repository   *Repository  `json:"repository,omitempty"`
	Sender       *User        `json:"sender,omitempty"`
	Installation *KeyValue    `json:"installation,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var p PullRequestEvent
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Action)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
