package ques01_test

import (
	"testing"
	"ttp/cmd/part4/ques01"
)

var dataProvider = []struct {
	noodle    ques01.Noodle
	soup      ques01.Soup
	topping   ques01.Topping
	seasoning ques01.Seasoning
	expected  int
}{
	{ques01.ThickNoodle, ques01.PorkBoneSoup, ques01.GreenOnion, ques01.Pepper, 700},
	{ques01.MiddleNoodle, ques01.SoySauceSoup, ques01.CharSiu, ques01.ShichimiPepper, 700},
	{ques01.ThickNoodle, ques01.SoySauceSoup, ques01.DriedSwaweed, ques01.Garlic, 700},
	{ques01.ThinNoodle, ques01.PorkBoneSoup, ques01.CharSiu, ques01.Garlic, 800},
	{ques01.VeryThinNoodle, ques01.MisoSoup, ques01.DriedSwaweed, ques01.ShichimiPepper, 600},
	{ques01.MiddleNoodle, ques01.SaltSoup, ques01.DriedSwaweed, ques01.Pepper, 600},
	{ques01.ThickNoodle, ques01.SaltSoup, ques01.BambooShoots, ques01.ShichimiPepper, 650},
	{ques01.MiddleNoodle, ques01.MisoSoup, ques01.GreenOnion, ques01.Garlic, 650},
	{ques01.VeryThinNoodle, ques01.MisoSoup, ques01.CharSiu, ques01.Pepper, 700},
	{ques01.VeryThinNoodle, ques01.SoySauceSoup, ques01.BambooShoots, ques01.Garlic, 650},
	{ques01.ThinNoodle, ques01.SaltSoup, ques01.GreenOnion, ques01.Pepper, 600},
	{ques01.ThinNoodle, ques01.SoySauceSoup, ques01.BambooShoots, ques01.Pepper, 600},
	{ques01.VeryThinNoodle, ques01.PorkBoneSoup, ques01.GreenOnion, ques01.ShichimiPepper, 650},
	{ques01.MiddleNoodle, ques01.PorkBoneSoup, ques01.BambooShoots, ques01.ShichimiPepper, 650},
	{ques01.ThinNoodle, ques01.MisoSoup, ques01.BambooShoots, ques01.ShichimiPepper, 600},
	{ques01.VeryThinNoodle, ques01.SaltSoup, ques01.CharSiu, ques01.Garlic, 750},
	{ques01.ThinNoodle, ques01.PorkBoneSoup, ques01.DriedSwaweed, ques01.Pepper, 650},
	{ques01.MiddleNoodle, ques01.SoySauceSoup, ques01.GreenOnion, ques01.ShichimiPepper, 600},
	{ques01.ThickNoodle, ques01.MisoSoup, ques01.CharSiu, ques01.ShichimiPepper, 750},
}

func TestGetPrice(t *testing.T) {
	for _, data := range dataProvider {
		result := ques01.GetPrice(data.noodle, data.soup, data.topping, data.seasoning)

		if result != data.expected {
			t.Fatalf("%s、%s、%s、%sの組み合わせが %d 円ではない。（実際： %d 円）",
				data.noodle,
				data.soup,
				data.topping,
				data.seasoning,
				data.expected,
				result)
		}
	}
}
