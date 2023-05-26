package ques06_test

import (
	"testing"
	"ttp/cmd/part1/ques06"
)

func TestEquipOneHandedWeaponWithShield(t *testing.T) {
	// Arrange
	character := ques06.NewCharacter(
		ques06.NewEmpty(),
		FactoryOneHandedShield())

	// Act
	result := character.EquipRightHand(FactoryOneHandedWeapon())

	// Assert
	if result != "" {
		t.Fatalf("片手に盾を持っている時に、片手武器が装備できない")
	}
}

func TestEquipTwoHandedWeaponWithShieldReturnError(t *testing.T) {
	// Arrange
	character := ques06.NewCharacter(
		ques06.NewEmpty(),
		FactoryOneHandedShield())

	// Act
	result := character.EquipRightHand(FactoryTwoHandedWeapon())

	// Assert
	if result != "装備できません" {
		t.Fatalf("片手に盾を持っている時に、両手武器が装備できてしまう")
	}
}

func FactoryOneHandedShield() ques06.Shield {
	return ques06.NewShield("木の盾", ques06.ONE_HANDED)
}

func FactoryOneHandedWeapon() ques06.Weapon {
	return ques06.NewWeapon("木の剣", ques06.ONE_HANDED)
}

func FactoryTwoHandedWeapon() ques06.Weapon {
	return ques06.NewWeapon("きこりの大斧", ques06.TWO_HANDED)
}
