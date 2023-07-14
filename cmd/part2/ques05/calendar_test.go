package ques05_test

import (
	"testing"
	"time"
	"ttp/cmd/part2/ques05"
)

var dataProvider = []struct {
	dateType string
	time     time.Time
	expected ques05.Color
}{
	{"祝日の日曜", factoryDate(2023, 1, 1), ques05.Red},
	{"祝日の土曜", factoryDate(2023, 2, 11), ques05.Red},
	{"祝日の平日", factoryDate(2023, 1, 9), ques05.Red},
	{"祝日ではない日曜", factoryDate(2023, 7, 2), ques05.Red},
	{"祝日ではない土曜", factoryDate(2023, 7, 1), ques05.Blue},
	{"祝日ではない平日", factoryDate(2023, 7, 3), ques05.Black},
}

func TestGetColor(t *testing.T) {
	for _, data := range dataProvider {
		result := ques05.GetColor(data.time)

		if result != data.expected {
			t.Fatalf(
				"日付が %s の時、文字色が %s ではない。 actual: %s",
				data.dateType,
				data.expected,
				result)
		}
	}
}

func factoryDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
