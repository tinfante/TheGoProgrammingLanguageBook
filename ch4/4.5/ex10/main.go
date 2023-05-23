package main

import (
	"fmt"
	"log"
	"os"
	"example.com/github"
    "sort"
    "time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
    sort.Slice(result.Items, func(i, j int) bool {
        return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt)
    })
	fmt.Printf("%d issues:\n", result.TotalCount)
    today := time.Now()
    monthAgo := today.AddDate(0, -1, 0)
    yearAgo := today.AddDate(-1, 0, 0)
    for _, item := range result.Items {
        var timeCategory string
        if item.CreatedAt.After(monthAgo) {
            timeCategory = "<1M"
        } else if item.CreatedAt.After(yearAgo) {
            timeCategory = "<1Y"
        } else {
            timeCategory = ">1Y"
        }
	    fmt.Printf("%s #%-5d %9.9s %.55s %v\n",
			timeCategory, item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
