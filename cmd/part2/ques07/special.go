package ques07

type Receive string

const (
	Shop     Receive = "店頭"
	Delivery Receive = "宅配"
)

type SpecialOffer string

const (
	FreeFrenchFry     SpecialOffer = "フライドポテト無料"
	FreeSecondPizza   SpecialOffer = "2枚目のピザ無料"
	Discount20Percent SpecialOffer = "20%割引"
)

func GetSpecialOffer(price int, receive Receive, hasCoupon bool) []SpecialOffer {
	result := []SpecialOffer{}

	if price >= 1500 {
		result = append(result, FreeFrenchFry)
	}

	if receive == Delivery {
		if hasCoupon {
			result = append(result, Discount20Percent)
		}
	} else if receive == Shop {
		result = append(result, FreeSecondPizza)
	}

	return result
}
