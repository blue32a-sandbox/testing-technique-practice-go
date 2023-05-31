package ques07

type Weapon struct {
	name        string
	power       int
	upgradeable bool
	isUpgraded  bool
}

func NewWeapon(name string, power int, upgradeable bool) Weapon {
	return Weapon{name: name, power: power, upgradeable: upgradeable, isUpgraded: false}
}

func (w Weapon) Power() int {
	return w.power
}

func (w Weapon) IsUpgraded() bool {
	return w.isUpgraded
}

func (w Weapon) Upgrade() (Weapon, string) {
	if w.upgradeable == false {
		return w, "武器の強化に失敗しました"
	}
	if w.isUpgraded == true {
		return w, "これ以上この武器の強化はできません"
	}

	enhanced := Weapon{
		name:        w.name,
		power:       w.power + 10,
		upgradeable: w.upgradeable,
		isUpgraded:  true}

	return enhanced, "武器の攻撃力が上がりました"
}
