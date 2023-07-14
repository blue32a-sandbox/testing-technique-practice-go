package ques02_test

import (
	"testing"
	"time"
	"ttp/cmd/part2/ques02"
)

var dataProvider = []struct {
	time            time.Time
	isSpecialMember bool
	expected        int
}{
	{FactoryBusinessDay(8, 45), true, 0},
	{FactoryBusinessDay(8, 45), false, 0},
	{FactoryBusinessDay(8, 44), true, 0},
	{FactoryBusinessDay(8, 44), false, 110},
	{FactoryBusinessHoliday(8, 45), true, 0},
	{FactoryBusinessHoliday(8, 45), false, 110},
	{FactoryBusinessHoliday(8, 44), true, 0},
	{FactoryBusinessHoliday(8, 44), false, 110},
}

func TestGetFee(t *testing.T) {
	for _, data := range dataProvider {
		result := ques02.GetFee(data.time, data.isSpecialMember)

		if result != data.expected {
			t.Fatalf(
				"日時が %s で特別会員が %t のとき、手数料が %d ではない。 actual: %d",
				data.time.Format("2006/1/2 15:04:00"),
				data.isSpecialMember,
				data.expected,
				result)
		}
	}
}

func FactoryBusinessHoliday(hour int, min int) time.Time {
	return time.Date(2023, 6, 3, hour, min, 0, 0, time.Local)
}

func FactoryBusinessDay(hour int, min int) time.Time {
	return time.Date(2023, 6, 2, hour, min, 0, 0, time.Local)
}
