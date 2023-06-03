package ques01_test

import (
	"testing"
	"time"
	"ttp/cmd/part2/ques01"
)

var dataProvider = []struct {
	useCoupon bool
	time      time.Time
	actual    int
}{
	{true, FactoryHappyHourTime(), 100},
	{true, FactoryNotHappyHourTime(), 100},
	{false, FactoryHappyHourTime(), 290},
	{false, FactoryNotHappyHourTime(), 490},
}

func TestGetBeerPrice(t *testing.T) {
	for _, data := range dataProvider {
		result := ques01.GetBeerPrice(data.useCoupon, data.time)

		if result != data.actual {
			t.Fatalf(
				"クーポン使用が %t で時間が %s のとき、料金が %d ではない。 actual: %d",
				data.useCoupon,
				data.time.Format("15:04"),
				data.actual,
				result)
		}
	}
}

func FactoryHappyHourTime() time.Time {
	return FactoryDate(16, 00)
}

func FactoryNotHappyHourTime() time.Time {
	return FactoryDate(15, 59)
}

func FactoryDate(hour int, min int) time.Time {
	return time.Date(2023, 6, 1, hour, min, 0, 0, time.Local)
}
