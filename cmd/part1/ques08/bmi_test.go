package ques08_test

import (
	"testing"
	"ttp/cmd/part1/ques08"
)

var bmiErrorDataProvider = []struct {
	value float64
}{
	{0.0},
	{100.0},
	{110.0},
}

func TestBMIValueError(t *testing.T) {
	for _, data := range bmiErrorDataProvider {
		_, err := ques08.NewBMI(data.value)

		if err == nil {
			t.Fatalf("BMIの値が %f のときエラーが発生しない。", data.value)
		}
	}
}

var bmiDataProvider = []struct {
	value  float64
	actual string
}{
	{1.0, "痩せ"},
	{15.0, "痩せ"},
	{18.4, "痩せ"},
	{18.5, "普通体重"},
	{22.0, "普通体重"},
	{24.9, "普通体重"},
	{25.0, "前肥満"},
	{27.0, "前肥満"},
	{29.9, "前肥満"},
	{30.0, "肥満（１度）"},
	{32.0, "肥満（１度）"},
	{34.9, "肥満（１度）"},
	{35.0, "肥満（２度）"},
	{37.0, "肥満（２度）"},
	{39.9, "肥満（２度）"},
	{40.0, "肥満（３度）"},
	{50.0, "肥満（３度）"},
	{99.9, "肥満（３度）"},
}

func TestGetNutritionalStatus(t *testing.T) {
	for _, data := range bmiDataProvider {
		bmi, _ := ques08.NewBMI(data.value)

		result := ques08.GetNutritionalStatus(bmi)

		if result != data.actual {
			t.Fatalf(
				"BMIの値が %.1f のとき栄養状態が %s ではない。 actual: %s",
				data.value,
				data.actual,
				result)
		}
	}
}
