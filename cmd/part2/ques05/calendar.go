package ques05

import "time"

type Color string

const (
	Red   = Color("赤")
	Blue  = Color("青")
	Black = Color("黒")
)

var holidays = [...]time.Time{
	time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),  // 祝日の日曜
	time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC),  // 祝日の平日
	time.Date(2023, 2, 11, 0, 0, 0, 0, time.UTC), // 祝日の土曜
}

func GetColor(t time.Time) Color {
	for _, holiday := range holidays {
		if t.Equal(holiday) {
			return Red
		}
	}
	if t.Weekday() == time.Sunday {
		return Red
	}
	if t.Weekday() == time.Saturday {
		return Blue
	}
	return Black
}
