package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	withinMonth := make([]*github.Issue, 0)
	withinYear := make([]*github.Issue, 0)
	overYear := make([]*github.Issue, 0)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(now.AddDate(0, -1, 0)):
			withinMonth = append(withinMonth, item)
		case item.CreatedAt.After(now.AddDate(-1, 0, 0)):
			withinYear = append(withinYear, item)
		default:
			overYear = append(overYear, item)
		}
	}

	fmt.Printf("--- 一ヶ月未満\n")
	printIssues(withinMonth)
	fmt.Printf("\n--- 一年未満\n")
	printIssues(withinYear)
	fmt.Printf("\n--- 一年以上\n")
	printIssues(overYear)
}

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
