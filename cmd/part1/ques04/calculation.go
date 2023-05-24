package ques04

import (
	"errors"
	"math"
)

func CalculationClothPrice(length float64) (int64, error) {
	length = math.Floor(length*10) / 10

	if length < 0.1 {
		return 0, errors.New("長さの下限値は 0.1m です")
	} else if length > 100 {
		return 0, errors.New("長さの上限値は 100.0m です")
	}

	var unitPrice int64

	if length > 3 {
		unitPrice = 350
	} else {
		unitPrice = 400
	}

	return int64(length * float64(unitPrice)), nil
}
