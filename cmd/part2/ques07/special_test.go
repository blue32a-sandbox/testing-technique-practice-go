package ques07_test

import (
	"slices"
	"testing"
	"ttp/cmd/part2/ques07"
)

var dataProvider = []struct {
	price     int
	receive   ques07.Receive
	hasCoupon bool
	expect    []ques07.SpecialOffer
}{
	{1500, ques07.Shop, false, []ques07.SpecialOffer{ques07.FreeFrenchFry, ques07.FreeSecondPizza}},
	{1500, ques07.Delivery, true, []ques07.SpecialOffer{ques07.FreeFrenchFry, ques07.Discount20Percent}},
	{1500, ques07.Delivery, false, []ques07.SpecialOffer{ques07.FreeFrenchFry}},
	{1499, ques07.Shop, false, []ques07.SpecialOffer{ques07.FreeSecondPizza}},
	{1499, ques07.Delivery, true, []ques07.SpecialOffer{ques07.Discount20Percent}},
	{1499, ques07.Delivery, false, []ques07.SpecialOffer{}},
}

func TestGetSpecialOffer(t *testing.T) {
	for _, data := range dataProvider {
		result := ques07.GetSpecialOffer(data.price, data.receive, data.hasCoupon)
		slices.Sort(result)

		slices.Sort(data.expect)

		if !slices.Equal(result, data.expect) {
			t.Errorf("特典がアンマッチ。 expected: %v, actual: %v", data.expect, result)
		}
	}
}
