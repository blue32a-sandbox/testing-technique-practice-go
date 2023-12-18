package ques01

type Noodle string

const (
	VeryThinNoodle Noodle = "極細麺"
	ThinNoodle     Noodle = "細麺"
	MiddleNoodle   Noodle = "中太麺"
	ThickNoodle    Noodle = "太麺"
)

type Soup string

const (
	SoySauceSoup Soup = "醤油"
	MisoSoup     Soup = "味噌"
	SaltSoup     Soup = "塩"
	PorkBoneSoup Soup = "トンコツ"
)

type Topping string

const (
	CharSiu      Topping = "チャーシュー"
	GreenOnion   Topping = "ねぎ"
	DriedSwaweed Topping = "海苔"
	BambooShoots Topping = "メンマ"
)

type Seasoning string

const (
	Pepper         Seasoning = "コショウ"
	ShichimiPepper Seasoning = "七味唐辛子"
	Garlic         Seasoning = "ニンニク"
)

func GetPrice(noodle Noodle, soup Soup, topping Topping, seasoning Seasoning) int {
	price := 600
	if noodle == ThickNoodle {
		price += 50
	}
	if soup == PorkBoneSoup {
		price += 50
	}
	if topping == CharSiu {
		price += 100
	}
	if seasoning == Garlic {
		price += 50
	}
	return price
}
