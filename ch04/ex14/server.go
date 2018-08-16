package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

// curl "https://api.github.com/repos/golang/go/issues"
// curl "https://api.github.com/repos/golang/go/milestones"
// curl "https://api.github.com/orgs/golang/members"

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var data htmlData
	if issues, err := getIssues(); err == nil {
		data.Issues = issues
	}
	if milestones, err := getMilestones(); err == nil {
		data.Milestones = milestones
	}
	if users, err := getMembers(); err == nil {
		data.Users = users
	}
	render(w, data)
}

type htmlData struct {
	Issues     []*Issue
	Milestones []*Milestone
	Users      []*User
}

var htmlTemplate = template.Must(template.New("htmlTemplate").Parse(`
<h1>issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Issues}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>

<h1>milestones</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Milestones}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>

<h1>members</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>User</th>
</tr>
{{range .Users}}
<tr>
  <td><img src='{{.AvatarURL}}' height='64px' width='64px'></td>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
</tr>
{{end}}
`))

func render(w io.Writer, data htmlData) {
	if err := htmlTemplate.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type Milestone struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Creator   *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login     string
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

func getIssues() ([]*Issue, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/golang/go/issues", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get issues: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func getMilestones() ([]*Milestone, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/golang/go/milestones", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get issues: %s", resp.Status)
	}

	var result []*Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func getMembers() ([]*User, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/golang/members", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get issues: %s", resp.Status)
	}

	var result []*User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
