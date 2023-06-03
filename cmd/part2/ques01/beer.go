package ques01

import "time"

func GetBeerPrice(useCoupon bool, t time.Time) int {
	if useCoupon {
		return 100
	}
	if t.Hour() >= 16 && t.Hour() <= 17 {
		return 290
	}
	return 490
}
