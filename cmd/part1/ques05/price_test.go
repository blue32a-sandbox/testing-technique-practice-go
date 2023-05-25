package ques05_test

import (
	"testing"
	"ttp/cmd/part1/ques05"
)

var dataProvider = []struct {
	age    int
	actual string
}{
	{0, "無料"},
	{2, "無料"},
	{5, "無料"},
	{6, "500円"},
	{9, "500円"},
	{12, "500円"},
	{13, "1,000円"},
	{15, "1,000円"},
	{17, "1,000円"},
	{18, "1,500円"},
	{50, "1,500円"},
	{130, "1,500円"},
}

func TestGetPrice(t *testing.T) {
	for _, data := range dataProvider {
		result, _ := ques05.GetPrice(data.age)

		if result != data.actual {
			t.Fatalf(
				"年齢が %d 歳のとき料金が %s ではない。 actual: %s",
				data.age,
				data.actual,
				result)
		}
	}
}

var errorDataProvider = []struct {
	age int
}{
	{-10},
	{-1},
	{131},
	{150},
}

func TestInvalidAge(t *testing.T) {
	for _, data := range errorDataProvider {
		_, err := ques05.GetPrice(data.age)

		if err == nil {
			t.Fatalf("年齢が %d 歳のときエラーが発生しない。", data.age)
		}
	}
}
