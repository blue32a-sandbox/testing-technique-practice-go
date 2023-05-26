package ques06

type EquipmentType int

const (
	EMPTY EquipmentType = iota
	ONE_HANDED
	TWO_HANDED
)

type Equipment interface {
	GetName() string
	GetType() EquipmentType
}

type Empty struct{}

func NewEmpty() Empty {
	return Empty{}
}

func (e Empty) GetName() string {
	return ""
}

func (e Empty) GetType() EquipmentType {
	return EMPTY
}

type Weapon struct {
	name          string
	equipmentType EquipmentType
}

func NewWeapon(name string, equipmentType EquipmentType) Weapon {
	return Weapon{name: name, equipmentType: equipmentType}
}

func (w Weapon) GetName() string {
	return w.name
}

func (w Weapon) GetType() EquipmentType {
	return w.equipmentType
}

type Shield struct {
	name          string
	equipmentType EquipmentType
}

func NewShield(name string, equipmentType EquipmentType) Shield {
	return Shield{name: name, equipmentType: equipmentType}
}

func (s Shield) GetName() string {
	return s.name
}

func (s Shield) GetType() EquipmentType {
	return s.equipmentType
}

type Character struct {
	rightHand Equipment
	leftHand  Equipment
}

func NewCharacter(rightHand Equipment, leftHand Equipment) Character {
	return Character{rightHand: rightHand, leftHand: leftHand}
}

func (c Character) GetRightHand() Equipment {
	return c.rightHand
}

func (c Character) GetLeftHand() Equipment {
	return c.leftHand
}

// 盾持ち状態でのテストをするための最低限の実装
func (c *Character) EquipRightHand(equipment Equipment) string {
	if equipment.GetType() == TWO_HANDED {
		return "装備できません"
	}
	c.rightHand = equipment
	return ""
}
