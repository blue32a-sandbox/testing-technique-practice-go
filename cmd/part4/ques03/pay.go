package ques03

import (
	"math"
)

type Ratio string

const (
	Ratio50_50 Ratio = "50:50"
	Ratio60_40 Ratio = "60:40"
	Ratio70_30 Ratio = "70:30"
	Ratio80_20 Ratio = "80:20"
	Ratio90_10 Ratio = "90:10"
	Ratio100_0 Ratio = "100:0"
)

var ourSideRatios = map[Ratio]float64{
	Ratio50_50: 50.0,
	Ratio60_40: 60.0,
	Ratio70_30: 70.0,
	Ratio80_20: 80.0,
	Ratio90_10: 90.0,
	Ratio100_0: 100.0,
}

type Result struct {
	ourSide   int
	otherSide int
	change    int
}

func NewResult(ourSide int, otherSide int, change int) Result {
	return Result{ourSide: ourSide, otherSide: otherSide, change: change}
}

func Calculation(ourSide int, otherSide int, amount int, ratio Ratio) Result {
	ourSideRatio := ourSideRatios[ratio]

	totalAmount := float64(amount)
	ourSideTotal := math.Ceil(totalAmount * (ourSideRatio / 100))
	otherSideTotal := totalAmount - ourSideTotal
	ourSidePerPerson := int(math.Floor(ourSideTotal / float64(ourSide)))
	otherSidePerPerson := int(math.Floor(otherSideTotal / float64(otherSide)))

	change := int(totalAmount) - (ourSidePerPerson*ourSide + otherSidePerPerson*otherSide)

	return NewResult(ourSidePerPerson, otherSidePerPerson, change)
}
