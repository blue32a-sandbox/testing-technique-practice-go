package ques04_test

import (
	"testing"
	"ttp/cmd/part1/ques04"
)

var dataProvider = []struct {
	legnth   float64
	expected int64
}{
	{0.1, 40},
	{1.5, 600},
	{3.0, 1200},
	{3.1, 1085},
	{50.5, 17675},
	{100.0, 35000},
}

func TestPurchase(t *testing.T) {
	for _, data := range dataProvider {

		result, _ := ques04.CalculationClothPrice(data.legnth)

		if result != data.expected {
			t.Fatalf(
				"購入する生地の長さが %.1f mのとき購入価格が %d 円ではない。 actual: %d",
				data.legnth,
				data.expected,
				result)
		}
	}
}

var errorDataProvider = []struct {
	legnth float64
}{
	{0.0},
	{100.1},
	{110.0},
}

func TestInvalidLength(t *testing.T) {
	for _, data := range errorDataProvider {
		_, err := ques04.CalculationClothPrice(data.legnth)

		if err == nil {
			t.Fatalf("畳数が %.1f mのときエラーが発生しない。", data.legnth)
		}
	}
}
