package github

import (
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

//因为ＪＳＯＮ成员名字和Go结构体成员名字并不相同，因此需要Ｇｏ语言结构体成员Ｔａｇ来
//指定对应的ＪＳＯＮ名字
type IssuesSearchResult struct {
	TotalCount int      `jason:"total_count"`
	Items      []*Issue // 为什么要持有指针
}

type Issue struct {
	Number    int
	HTMLURL   string `json:html_url`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
