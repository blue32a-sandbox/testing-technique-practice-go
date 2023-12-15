package ques08_test

import (
	"testing"
	"ttp/cmd/part2/ques08"
)

var dataProvider = []struct {
	user     ques08.User
	day      int
	expected int
}{
	{ques08.NewUser(true, true, true), 1, 10},
	{ques08.NewUser(true, true, false), 1, 5},
	{ques08.NewUser(true, false, true), 1, 8},
	{ques08.NewUser(true, false, false), 1, 3},
	{ques08.NewUser(false, true, true), 1, 7},
	{ques08.NewUser(false, true, false), 1, 2},
	{ques08.NewUser(false, false, true), 1, 5},
	{ques08.NewUser(false, false, false), 1, 0},

	{ques08.NewUser(true, true, true), 15, 13},
	{ques08.NewUser(true, true, false), 15, 8},
	{ques08.NewUser(true, false, true), 15, 11},
	{ques08.NewUser(true, false, false), 15, 6},
	{ques08.NewUser(false, true, true), 15, 10},
	{ques08.NewUser(false, true, false), 15, 5},
	{ques08.NewUser(false, false, true), 15, 8},
	{ques08.NewUser(false, false, false), 15, 3},
}

func TestGetPointMultiplier(t *testing.T) {
	for _, data := range dataProvider {
		result := ques08.GetPointMultiplier(data.user, data.day)
		if result != data.expected {
			t.Fatalf("ユーザーが %#v で日付が %d のとき、ポイント倍率が %d ではない。（実際： %d）",
				data.user,
				data.day,
				data.expected,
				result)
		}
	}
}
