package ques01

import "math"

type Thermometer struct {
	temperature float64
}

func NewThermometer(temperature float64) Thermometer {
	return Thermometer{math.Ceil(temperature*10) / 10}
}

func (t Thermometer) Message() string {
	if t.temperature < 24.0 {
		return "寒い"
	} else if t.temperature < 26.0 {
		return "快適"
	}
	return "暑い"
}
