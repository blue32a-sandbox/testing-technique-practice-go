package ques06_test

import (
	"testing"
	"ttp/cmd/part2/ques06"
)

var funcRaffleWin = func() bool { return true }
var funcRaffleNotWin = func() bool { return false }

var dataProvider = []struct {
	description         string
	rank                ques06.MembershipRank
	charge              ques06.ChargeAmount
	funcRaffle          ques06.FuncRaffle
	expectedPlusBonus   ques06.PlusBonus
	expectedIsGetCoupon bool
}{
	{"シルバー会員が3,000円チャージ", ques06.Silver, ques06.Charge3000, funcRaffleWin, ques06.Plus1, false},
	{"シルバー会員が5,000円チャージ、クーポン抽選は当選", ques06.Silver, ques06.Charge5000, funcRaffleWin, ques06.Plus2, true},
	{"シルバー会員が5,000円チャージ、クーポン抽選は外れ", ques06.Silver, ques06.Charge5000, funcRaffleNotWin, ques06.Plus2, false},
	{"シルバー会員が10,000円チャージ、クーポン抽選は当選", ques06.Silver, ques06.Charge10000, funcRaffleWin, ques06.Plus4, true},
	{"シルバー会員が10,000円チャージ、クーポン抽選は外れ", ques06.Silver, ques06.Charge10000, funcRaffleNotWin, ques06.Plus4, false},
	{"ゴールド会員が3,000円チャージ", ques06.Gold, ques06.Charge3000, funcRaffleWin, ques06.Plus3, false},
	{"ゴールド会員が5,000円チャージ、クーポン抽選は当選", ques06.Gold, ques06.Charge5000, funcRaffleWin, ques06.Plus5, true},
	{"ゴールド会員が5,000円チャージ、クーポン抽選は外れ", ques06.Gold, ques06.Charge5000, funcRaffleNotWin, ques06.Plus5, false},
	{"ゴールド会員が10,000円チャージ、クーポン抽選は当選", ques06.Gold, ques06.Charge10000, funcRaffleWin, ques06.Plus10, true},
	{"ゴールド会員が10,000円チャージ、クーポン抽選は外れ", ques06.Gold, ques06.Charge10000, funcRaffleNotWin, ques06.Plus10, false},
	{"ブラック会員が3,000円チャージ", ques06.Black, ques06.Charge3000, funcRaffleWin, ques06.Plus5, false},
	{"ブラック会員が5,000円チャージ、クーポン抽選は当選", ques06.Black, ques06.Charge5000, funcRaffleWin, ques06.Plus7, true},
	{"ブラック会員が5,000円チャージ、クーポン抽選は外れ", ques06.Black, ques06.Charge5000, funcRaffleNotWin, ques06.Plus7, false},
	{"ブラック会員が10,000円チャージ、クーポン抽選は当選", ques06.Black, ques06.Charge10000, funcRaffleWin, ques06.Plus15, true},
	{"ブラック会員が10,000円チャージ、クーポン抽選は外れ", ques06.Black, ques06.Charge10000, funcRaffleNotWin, ques06.Plus15, false},
}

func TestGetSpecialOffer(t *testing.T) {
	for _, data := range dataProvider {
		t.Run(data.description, func(t *testing.T) {
			plusBonus, isGetCoupon := ques06.GetSpecialOffer(data.rank, data.charge, data.funcRaffle)
			if plusBonus != data.expectedPlusBonus {
				t.Errorf("会員ランクが %s かつチャージ額が %d円 で、ボーナスが %d%%ではない。actual: %d%%",
					data.rank,
					data.charge,
					data.expectedPlusBonus,
					plusBonus)
			}
			if isGetCoupon != data.expectedIsGetCoupon {
				t.Errorf("会員ランクが %s かつチャージ額が %d円 で、クーポン抽選結果が %t ではない。actual: %t",
					data.rank,
					data.charge,
					data.expectedIsGetCoupon,
					isGetCoupon)
			}
		})
	}
}
