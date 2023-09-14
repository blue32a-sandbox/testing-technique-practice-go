package ques09

import "time"

func Rate(date time.Time) int {
	if date.Day() >= 1 && date.Day() <= 5 {
		return 20
	}
	if date.Day() >= 28 && date.Day() <= 30 {
		return 20
	}
	return 0
}
