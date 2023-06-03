package ques02

import "time"

func GetFee(time time.Time, isSpecialMember bool) int {
	if isSpecialMember {
		return 0
	}
	if isBusinessHoliday(time) {
		return 110
	}
	if ((time.Hour() == 8 && time.Minute() >= 45) || time.Hour() > 8) &&
		time.Hour() <= 17 {
		return 0
	}
	return 110
}

func isBusinessHoliday(t time.Time) bool {
	// 簡略化のため、祝日は判定しない
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}
