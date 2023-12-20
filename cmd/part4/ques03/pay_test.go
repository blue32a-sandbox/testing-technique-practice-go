package ques03_test

import (
	"testing"
	"ttp/cmd/part4/ques03"
)

var dataProvider = []struct {
	ourSide   int
	otherSide int
	amount    int
	ratio     ques03.Ratio
	expected  ques03.Result
}{
	{50, 1, 1, ques03.Ratio100_0, ques03.NewResult(0, 0, 1)},
	{99, 50, 1, ques03.Ratio90_10, ques03.NewResult(0, 0, 1)},
	{1, 1, 500000, ques03.Ratio80_20, ques03.NewResult(400000, 100000, 0)},
	{99, 99, 999999, ques03.Ratio80_20, ques03.NewResult(8080, 2020, 99)},
	{1, 99, 1, ques03.Ratio70_30, ques03.NewResult(1, 0, 0)},
	{50, 50, 999999, ques03.Ratio80_20, ques03.NewResult(16000, 3999, 49)},
	{50, 99, 500000, ques03.Ratio50_50, ques03.NewResult(5000, 2525, 25)},
	{99, 50, 500000, ques03.Ratio100_0, ques03.NewResult(5050, 0, 50)},
	{1, 1, 999999, ques03.Ratio60_40, ques03.NewResult(600000, 399999, 0)},
	{1, 1, 500000, ques03.Ratio90_10, ques03.NewResult(450000, 50000, 0)},
	{1, 99, 999999, ques03.Ratio100_0, ques03.NewResult(999999, 0, 0)},
	{1, 50, 500000, ques03.Ratio60_40, ques03.NewResult(300000, 4000, 0)},
	{99, 99, 1, ques03.Ratio60_40, ques03.NewResult(0, 0, 1)},
	{99, 1, 999999, ques03.Ratio50_50, ques03.NewResult(5050, 499999, 50)},
	{50, 99, 999999, ques03.Ratio90_10, ques03.NewResult(18000, 1010, 9)},
	{99, 50, 500000, ques03.Ratio70_30, ques03.NewResult(3535, 3000, 35)},
	{1, 50, 1, ques03.Ratio50_50, ques03.NewResult(1, 0, 0)},
	{50, 99, 1, ques03.Ratio60_40, ques03.NewResult(0, 0, 1)},
	{99, 1, 1, ques03.Ratio80_20, ques03.NewResult(0, 0, 1)},
	{50, 1, 999999, ques03.Ratio70_30, ques03.NewResult(14000, 299999, 0)},
}

func TestCalculation(t *testing.T) {
	for _, data := range dataProvider {
		result := ques03.Calculation(data.ourSide, data.otherSide, data.amount, data.ratio)

		if result != data.expected {
			t.Fatalf("自分側の人数: %d, 相手側の人数: %d, 金額: %d, 割合: %s, 期待: %#v, 実際: %#v",
				data.ourSide,
				data.otherSide,
				data.amount,
				data.ratio,
				data.expected,
				result)
		}
	}

}
