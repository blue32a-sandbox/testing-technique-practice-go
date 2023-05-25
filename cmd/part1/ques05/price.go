package ques05

import "errors"

func GetPrice(age int) (string, error) {
	if age < 0 || age > 130 {
		return "", errors.New("年齢は0から130までで入力してください")
	}

	if age < 6 {
		return "無料", nil
	} else if age < 13 {
		return "500円", nil
	} else if age < 18 {
		return "1,000円", nil
	}
	return "1,500円", nil
}
