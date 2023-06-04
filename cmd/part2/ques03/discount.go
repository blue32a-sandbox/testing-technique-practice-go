package ques03

type ItemCategory int

const (
	WAI_SHIRT ItemCategory = iota
	TIE
	OTHER
)

type Item struct {
	category ItemCategory
}

func NewItem(category ItemCategory) Item {
	return Item{category: category}
}

func GetDiscountRate(items []Item) int {
	rate := 0
	if len(items) >= 7 {
		rate += 7
	}
	hasWaiShirt := false
	for _, item := range items {
		if item.category == WAI_SHIRT {
			hasWaiShirt = true
			break
		}
	}
	if hasWaiShirt {
		for _, item := range items {
			if item.category == TIE {
				rate += 5
				break
			}
		}
	}
	return rate
}
