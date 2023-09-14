package ques09_test

import (
	"testing"
	"time"
	"ttp/cmd/part1/ques09"
)

var dataProvider = []struct {
	date     time.Time
	expected int
}{
	{factoryDate(2023, 8, 1), 20},
	{factoryDate(2023, 8, 3), 20},
	{factoryDate(2023, 8, 5), 20},
	{factoryDate(2023, 8, 6), 0},
	{factoryDate(2023, 8, 15), 0},
	{factoryDate(2023, 8, 27), 0},
	{factoryDate(2023, 8, 28), 20},
	{factoryDate(2023, 8, 29), 20},
	{factoryDate(2023, 8, 30), 20},
	{factoryDate(2023, 8, 31), 0},
}

func TestRate(t *testing.T) {
	for _, data := range dataProvider {
		result := ques09.Rate(data.date)

		if result != data.expected {
			t.Fatalf("日付が %s 日の時、割引率が %d %%ではない",
				data.date.Format("2"),
				data.expected)
		}
	}

}

func factoryDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
