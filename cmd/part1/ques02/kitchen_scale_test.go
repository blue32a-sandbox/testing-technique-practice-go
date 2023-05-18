package ques02_test

import (
	"testing"
	"ttp/cmd/part1/ques02"
)

var dataProvider = []struct {
	weight int64
	actual string
}{
	{0, "0g"},
	{1000, "1000g"},
	{2000, "2000g"},
	{2001, "EEEE"},
	{2010, "EEEE"},
}

func TestPlusMass(t *testing.T) {
	for _, data := range dataProvider {
		object, _ := ques02.NewObject(data.weight)
		kitchenScale := ques02.NewKitchenScale().
			PowerOn().
			PutOn(object)

		result := kitchenScale.ViewTotalMass()

		if result != data.actual {
			t.Fatalf(
				"重さが %d のとき表示メッセージが %s ではない。 actual: %s",
				kitchenScale.TotalMass(),
				data.actual,
				result)
		}
	}
}

var minusDataProvider = []struct {
	weight int64
	actual string
}{
	{1, "EEEE"},
	{10, "EEEE"},
}

func TestMinusMass(t *testing.T) {
	for _, data := range minusDataProvider {
		object, _ := ques02.NewObject(data.weight)
		kitchenScale := ques02.NewKitchenScale().
			PutOn(object).
			PowerOn().
			TakeOff()

		result := kitchenScale.ViewTotalMass()

		if result != data.actual {
			t.Fatalf(
				"重さが %d のとき表示メッセージが %s ではない。 actual: %s",
				kitchenScale.TotalMass(),
				data.actual,
				result)
		}
	}
}
