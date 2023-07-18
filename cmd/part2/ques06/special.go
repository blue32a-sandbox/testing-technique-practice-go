package ques06

type MembershipRank string

const (
	Silver MembershipRank = "シルバー"
	Gold   MembershipRank = "ゴールド"
	Black  MembershipRank = "ブラック"
)

type ChargeAmount int

const (
	Charge3000  ChargeAmount = 3000
	Charge5000  ChargeAmount = 5000
	Charge10000 ChargeAmount = 10000
)

type PlusBonus int

const (
	Plus1  PlusBonus = 1
	Plus2  PlusBonus = 2
	Plus3  PlusBonus = 3
	Plus4  PlusBonus = 4
	Plus5  PlusBonus = 5
	Plus7  PlusBonus = 7
	Plus10 PlusBonus = 10
	Plus15 PlusBonus = 15
)

type FuncRaffle func() bool

var bonusMap = map[MembershipRank]map[ChargeAmount]PlusBonus{
	Silver: {
		Charge3000:  Plus1,
		Charge5000:  Plus2,
		Charge10000: Plus4,
	},
	Gold: {
		Charge3000:  Plus3,
		Charge5000:  Plus5,
		Charge10000: Plus10,
	},
	Black: {
		Charge3000:  Plus5,
		Charge5000:  Plus7,
		Charge10000: Plus15,
	},
}

func GetSpecialOffer(rank MembershipRank, charge ChargeAmount, funcRaffle FuncRaffle) (PlusBonus, bool) {
	plusBonus := Plus1
	if bonus, ok := bonusMap[rank][charge]; ok {
		plusBonus = bonus
	}

	isGetCoupon := false
	if charge != Charge3000 {
		isGetCoupon = funcRaffle()
	}

	return plusBonus, isGetCoupon
}
