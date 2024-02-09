package ques05_test

import (
	"testing"
	"ttp/cmd/part4/ques05"
)

var dataProvider = []struct {
	gender ques05.Gender
	age    int
	option ques05.Option
}{
	{ques05.Male, 44, ques05.None},
	{ques05.Male, 60, ques05.MetabolicSyndrome},
	{ques05.Female, 45, ques05.MetabolicSyndrome},
	{ques05.Female, 60, ques05.None},
	{ques05.Female, 21, ques05.Fundus},
	{ques05.Male, 34, ques05.Fundus},
	{ques05.Female, 45, ques05.Gynecology},
	{ques05.Female, 44, ques05.Gynecology},
	{ques05.Male, 46, ques05.MetabolicSyndrome},
	{ques05.Female, 35, ques05.Gynecology},
	{ques05.Male, 45, ques05.None},
	{ques05.Male, 35, ques05.MetabolicSyndrome},
	{ques05.Female, 44, ques05.MetabolicSyndrome},
	{ques05.Female, 60, ques05.Gynecology},
	{ques05.Female, 36, ques05.Gynecology},
	{ques05.Male, 37, ques05.None},
	{ques05.Male, 38, ques05.MetabolicSyndrome},
	{ques05.Male, 20, ques05.None},
	{ques05.Female, 47, ques05.Brain},
	{ques05.Female, 60, ques05.Fundus},
	{ques05.Female, 48, ques05.Gynecology},
	{ques05.Male, 49, ques05.Fundus},
	{ques05.Female, 20, ques05.Fundus},
	{ques05.Male, 50, ques05.None},
	{ques05.Female, 39, ques05.Fundus},
	{ques05.Female, 44, ques05.Fundus},
	{ques05.Male, 22, ques05.None},
	{ques05.Male, 45, ques05.Fundus},
	{ques05.Female, 34, ques05.None},
	{ques05.Female, 35, ques05.Fundus},
	{ques05.Female, 35, ques05.None},
	{ques05.Male, 45, ques05.Brain},
	{ques05.Female, 60, ques05.Brain},
}

func TestRequest(t *testing.T) {
	for _, data := range dataProvider {
		_, err := ques05.Request(data.gender, data.age, data.option)

		if err != nil {
			t.Fatalf("%s %d %sでエラーが発生しました: %s", data.gender, data.age, data.option, err)
		}
	}
}
