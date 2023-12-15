package ques10_test

import (
	"testing"
	"ttp/cmd/part1/ques10"
)

var dataProvider = []struct {
	day      int
	expected int
}{
	{1, 20},
	{3, 20},
	{5, 20},
	{6, 0},
	{15, 0},
	{27, 0},
	{28, 20},
	{29, 20},
	{30, 20},
	{31, 0},
}

func TestGetDiscountRate(t *testing.T) {
	for _, data := range dataProvider {
		result := ques10.GetDiscountRate(data.day)
		if result != data.expected {
			t.Fatalf("日付が %d のとき、割引率が %d ではない。（実際： %d）",
				data.day,
				data.expected,
				result)
		}
	}
}
