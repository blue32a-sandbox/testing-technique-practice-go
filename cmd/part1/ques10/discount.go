package ques10

func GetDiscountRate(day int) int {
	if day >= 1 && day <= 5 {
		return 20
	} else if day >= 28 && day <= 30 {
		return 20
	}
	return 0
}
