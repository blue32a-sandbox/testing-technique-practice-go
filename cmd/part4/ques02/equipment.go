package ques02

type Sword string

const (
	CypressSword Sword = "ひのきのぼう"
	IronSword    Sword = "鉄の剣"
	CarrotSword  Sword = "ニンジンソード"
)

var swordPrices = map[Sword]int{
	CypressSword: 50,
	IronSword:    100,
	CarrotSword:  300,
}

type Armor string

const (
	ClothArmor         Armor = "ぬののふく"
	IronArmor          Armor = "鉄の鎧"
	MascotCostumeArmor Armor = "きぐるみ"
)

var armorPrices = map[Armor]int{
	ClothArmor:         50,
	IronArmor:          100,
	MascotCostumeArmor: 250,
}

type Shield string

const (
	LeatherShield     Shield = "かわのたて"
	IronShield        Shield = "鉄の盾"
	TurtleShellShield Shield = "亀の甲羅"
)

var shieldPrices = map[Shield]int{
	LeatherShield:     50,
	IronShield:        100,
	TurtleShellShield: 200,
}

func GetPrice(sword Sword, armor Armor, shield Shield) int {
	price := 0
	price += swordPrices[sword]
	price += armorPrices[armor]
	price += shieldPrices[shield]
	return price
}
