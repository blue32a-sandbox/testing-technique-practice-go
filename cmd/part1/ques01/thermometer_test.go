package ques01_test

import (
	"testing"
	"ttp/cmd/part1/ques01"
)

var dataProvider = []struct {
	temperature float64
	actual      string
}{
	{20.0, "寒い"},
	{23.9, "寒い"},
	{24.0, "快適"},
	{25.0, "快適"},
	{25.9, "快適"},
	{26.0, "暑い"},
	{30.0, "暑い"},
}

func TestMessage(t *testing.T) {
	for _, data := range dataProvider {
		thermometer := ques01.NewThermometer(data.temperature)

		result := thermometer.Message()

		if result != data.actual {
			t.Fatalf(
				"室温が %.1f のとき表示メッセージが %s ではない。 actual: %s",
				data.temperature,
				data.actual, result)
		}
	}
}
