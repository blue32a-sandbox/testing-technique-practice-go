package ques05

import "errors"

type Gender string

const (
	Male   Gender = "男性"
	Female Gender = "女性"
)

type Option string

const (
	None              Option = "なし"
	Fundus            Option = "眼底検診"
	MetabolicSyndrome Option = "メタボ検診"
	Gynecology        Option = "婦人科検診"
	Brain             Option = "脳検診"
)

func Request(gender Gender, age int, option Option) (string, error) {
	if (option == MetabolicSyndrome || option == Gynecology) && age < 35 {
		return "", errors.New("メタボ健診と婦人科検診は35歳以上からです")
	}
	if option == Brain && age < 45 {
		return "", errors.New("脳検診は45歳以上からです")
	}
	if option == Gynecology && gender != Female {
		return "", errors.New("婦人科検診は女性のみです")
	}

	return "OK", nil
}
