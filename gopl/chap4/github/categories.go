package github

import (
	"time"
)

type TimeStamp int

const (
	LessThanOneMonth TimeStamp = iota
	MoreThanOneMonthLessThanOneYear
	MoreThanOneYear
)

func GetTimeCate(t time.Time) TimeStamp {
	oneMonthLater := t.AddDate(0, 1, 0)
	oneYearLater := t.AddDate(1, 0, 0)

	now := time.Now()
	if now.Before(oneMonthLater) {
		return LessThanOneMonth
	}

	if now.After(oneYearLater) {
		return MoreThanOneYear
	}

	return MoreThanOneMonthLessThanOneYear
}
