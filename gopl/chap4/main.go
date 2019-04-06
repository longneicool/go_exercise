// this is for 4-10
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go_exercise/gopl/chap4/github"
)

func main() {
	addr := flag.String("addr", "", "The repo addr")
	timeType := flag.Int("t", 0, "The category of issue printed")

	flag.Parse()

	urls := strings.Split(*addr, " ")
	fmt.Println(urls)
	result, err := github.SearchIssues(urls)
	if err != nil {
		log.Fatal(err)
	}

	issuesByTime := make(map[github.TimeStamp]string)
	for _, item := range result.Items {
		issue := fmt.Sprintf("#%-5d %9.9s %.55s %.55s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(time.UnixDate))

		cate := github.GetTimeCate(item.CreatedAt)
		issuesByTime[cate] += issue
	}

	fmt.Println(issuesByTime[github.TimeStamp(*timeType)])
}
