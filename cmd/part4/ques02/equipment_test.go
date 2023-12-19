package ques02_test

import (
	"testing"
	"ttp/cmd/part4/ques02"
)

var dataProvider = []struct {
	sword    ques02.Sword
	armor    ques02.Armor
	shield   ques02.Shield
	expected int
}{
	{ques02.IronSword, ques02.MascotCostumeArmor, ques02.LeatherShield, 400},
	{ques02.CarrotSword, ques02.ClothArmor, ques02.TurtleShellShield, 550},
	{ques02.CypressSword, ques02.IronArmor, ques02.LeatherShield, 200},
	{ques02.CarrotSword, ques02.MascotCostumeArmor, ques02.IronShield, 650},
	{ques02.CypressSword, ques02.ClothArmor, ques02.IronShield, 200},
	{ques02.IronSword, ques02.IronArmor, ques02.TurtleShellShield, 400},
	{ques02.CarrotSword, ques02.ClothArmor, ques02.LeatherShield, 400},
	{ques02.CypressSword, ques02.MascotCostumeArmor, ques02.TurtleShellShield, 500},
	{ques02.IronSword, ques02.ClothArmor, ques02.IronShield, 250},
	{ques02.CarrotSword, ques02.IronArmor, ques02.IronShield, 500},
}

func TestGetPrice(t *testing.T) {
	for _, data := range dataProvider {
		result := ques02.GetPrice(data.sword, data.armor, data.shield)

		if result != data.expected {
			t.Fatalf("%s、%s、%sの組み合わせが %d 円ではない。（実際： %d 円）",
				data.sword,
				data.armor,
				data.shield,
				data.expected,
				result)
		}
	}

}
