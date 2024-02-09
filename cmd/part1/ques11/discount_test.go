package ques11_test

import (
	"testing"
	"ttp/cmd/part1/ques11"
)

var dataProvider = []struct {
	day      int
	expected int
}{
	{1, 20},
	{2, 20},
	{3, 20},
	{4, 20},
	{5, 20},
	{28, 20},
	{29, 20},
	{30, 20},
	{15, 0}, // 上記以外の代表値
}

func TestGetDiscountRate(t *testing.T) {
	for _, data := range dataProvider {
		result := ques11.GetDiscountRate(data.day)
		if result != data.expected {
			t.Fatalf("日付が %d のとき、割引率が %d ではない。（実際： %d）",
				data.day,
				data.expected,
				result)
		}
	}
}
