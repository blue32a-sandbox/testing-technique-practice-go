package ques07_test

import (
	"testing"
	"ttp/cmd/part1/ques07"
)

func TestUpgradeUpgradeableWeapon(t *testing.T) {
	wepon := FactoryUpgradeableWepon(50)

	upgraded, message := wepon.Upgrade()

	if message != "武器の攻撃力が上がりました" {
		t.Fatalf("メッセージが 武器の攻撃力が上がりました ではない。 actual: %s", message)
	}
	if upgraded.Power() != 60 {
		t.Fatalf("攻撃力が10増えた 60 ではない。 actual: %d", upgraded.Power())
	}
	if upgraded.IsUpgraded() != true {
		t.Fatalf("強化済みになっていない。 actual: %t", upgraded.IsUpgraded())
	}
}

func TestUpgradeAlreadyUpgradedWeapon(t *testing.T) {
	wepon := FactoryUpgradeableWepon(50)

	upgraded, _ := wepon.Upgrade()
	reUpgraded, message := upgraded.Upgrade()

	if message != "これ以上この武器の強化はできません" {
		t.Fatalf("メッセージが これ以上この武器の強化はできません ではない。 actual: %s", message)
	}
	if reUpgraded.Power() != 60 {
		t.Fatalf("攻撃力が一度の強化で10増えた 60 ではない。 actual: %d", reUpgraded.Power())
	}
}

func TestUpgradeNonUpgradeableWeapon(t *testing.T) {
	wepon := FactoryNonUpgradeableWepon(5)

	upgraded, message := wepon.Upgrade()

	if message != "武器の強化に失敗しました" {
		t.Fatalf("メッセージが 武器の強化に失敗しました ではない。 actual: %s", message)
	}
	if upgraded.Power() != 5 {
		t.Fatalf("攻撃力がもともとの 5 ではない。 actual: %d", upgraded.Power())
	}
	if upgraded.IsUpgraded() != false {
		t.Fatalf("強化済みになっている。 actual: %t", upgraded.IsUpgraded())
	}
}

func FactoryUpgradeableWepon(power int) ques07.Weapon {
	return ques07.NewWeapon("龍の槍", power, true)
}

func FactoryNonUpgradeableWepon(power int) ques07.Weapon {
	return ques07.NewWeapon("けやきのぼう", power, false)
}
