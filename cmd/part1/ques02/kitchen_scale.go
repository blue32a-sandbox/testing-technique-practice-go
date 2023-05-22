package ques02

import (
	"errors"
	"fmt"
)

type Object struct {
	weight int64
}

func NewObject(weight int64) (Object, error) {
	if weight < 0 {
		return Object{}, errors.New("The weight must be greater than or equal to 0.")
	}
	return Object{weight}, nil
}

type KitchenScale struct {
	isPowerOn bool
	totalMass int64
	objects   []Object
}

func NewKitchenScale() KitchenScale {
	return KitchenScale{false, 0, []Object{}}
}

func (ks KitchenScale) TotalMass() int64 {
	return ks.totalMass
}

func (ks KitchenScale) PowerOn() KitchenScale {
	return KitchenScale{true, 0, ks.objects}
}

func (ks KitchenScale) PowerOff() KitchenScale {
	return KitchenScale{false, ks.totalMass, ks.objects}
}

func (ks KitchenScale) PutOn(object Object) KitchenScale {
	totalMass := ks.totalMass
	if ks.isPowerOn {
		totalMass += object.weight
	}
	return KitchenScale{ks.isPowerOn, totalMass, append(ks.objects, object)}
}

func (ks KitchenScale) TakeOff() KitchenScale {
	totalMass := ks.totalMass
	if ks.isPowerOn {
		totalMass -= ks.objects[len(ks.objects)-1].weight
	}
	return KitchenScale{ks.isPowerOn, totalMass, ks.objects[:len(ks.objects)-1]}
}

func (ks KitchenScale) ViewTotalMass() string {
	if ks.totalMass < 0 || ks.totalMass > 2000 {
		return "EEEE"
	}
	return fmt.Sprintf("%dg", ks.totalMass)
}
