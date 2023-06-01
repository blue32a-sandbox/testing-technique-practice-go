package ques08

import "errors"

type BMI struct {
	value float64
}

func NewBMI(value float64) (BMI, error) {
	if value <= 0 || value >= 100 {
		return BMI{}, errors.New("1以上かつ99以下の値を入力してください")
	}

	filteredValue := float64(int(value*10)) / 10

	return BMI{value: filteredValue}, nil
}

func GetNutritionalStatus(bmi BMI) string {
	if bmi.value < 18.5 {
		return "痩せ"
	} else if bmi.value <= 24.9 {
		return "普通体重"
	} else if bmi.value <= 29.9 {
		return "前肥満"
	} else if bmi.value <= 34.9 {
		return "肥満（１度）"
	} else if bmi.value <= 39.9 {
		return "肥満（２度）"
	}
	return "肥満（３度）"
}
