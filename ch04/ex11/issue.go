package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// todo
// - [x] create issue
// - [x] list issues
// - [x] show issue
// - [ ] update issue
// - [x] close issue

func main() {
	action := flag.String("action", "list", "action")
	issueNumber := flag.Int("number", 0, "issue number")
	issueTitle := flag.String("title", "", "issue number")
	issueBody := flag.String("body", "", "issue number")
	flag.Parse()

	switch *action {
	case "create":
		createIssue(*issueTitle, *issueBody)
	case "list":
		showIssues()
	case "show":
		showIssue(*issueNumber)
		showIssueComments(*issueNumber)
	case "update":
		fmt.Printf("todo update %d", *issueNumber)
	case "close":
		closeIssue(*issueNumber)
	default:
		fmt.Fprintf(os.Stderr, "bye!\n")
		os.Exit(1)
	}
}

type Issue struct {
	Number    int       `json:",omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:"title"`
	State     string    `json:",omitempty"`
	User      *User     `json:",omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body,omitempty"`
}

type Comment struct {
	HTMLURL   string `json:"html_url"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login     string
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

func closeIssue(issueNumber int) {
	url := fmt.Sprintf("https://api.github.com/repos/iinm/hello-gopl/issues/%d", issueNumber)
	req, err := http.NewRequest(
		"PATCH",
		url,
		bytes.NewBuffer([]byte(`{"state": "closed"}`)),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "token "+os.Getenv("github_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed to patch issues: %s\n", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Printf("url: %s\ntitle: %s\nstate: %s\nbody: %s\n",
		result.HTMLURL, result.Title, result.State, result.Body)
}

func createIssue(title, body string) {
	issue := Issue{Title: title, Body: body}
	jsonBytes, _ := json.Marshal(issue)
	fmt.Println(string(jsonBytes))

	req, err := http.NewRequest(
		"POST",
		"https://api.github.com/repos/iinm/hello-gopl/issues",
		bytes.NewBuffer([]byte(jsonBytes)),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "token "+os.Getenv("github_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		log.Fatalf("failed to post issues: %s\n", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Printf("url: %s\ntitle: %s\nstate: %s\nbody: %s\n",
		result.HTMLURL, result.Title, result.State, result.Body)
}

func showIssues() {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/iinm/hello-gopl/issues", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed to get issues: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()

	for _, item := range result {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func showIssue(issueNumber int) {
	url := fmt.Sprintf("https://api.github.com/repos/iinm/hello-gopl/issues/%d", issueNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed to get issues: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Printf("url: %s\ntitle: %s\nstate: %s\nbody: %s\n",
		result.HTMLURL, result.Title, result.State, result.Body)
}

func showIssueComments(issueNumber int) {
	url := fmt.Sprintf("https://api.github.com/repos/iinm/hello-gopl/issues/%d/comments", issueNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("failed to get issues: %s", resp.Status)
	}

	var result []*Comment
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()

	for _, item := range result {
		fmt.Printf("--- %s\n[%s] %s:\n%s\n",
			item.HTMLURL, item.CreatedAt, item.User.Login, item.Body)
	}
}
