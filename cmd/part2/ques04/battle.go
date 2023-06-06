package ques04

type PartyMember struct {
	job string
}

func NewPartyMember(job string) PartyMember {
	return PartyMember{job: job}
}

type Party []PartyMember

type Item struct {
	name string
}

func NewItem(name string) Item {
	return Item{name: name}
}

type Items []Item

// 簡略化して最後の進行状態を返す
func GetProgress(party Party, items Items) string {
	isSageInParty := false
	for _, member := range party {
		if member.job == "賢者" {
			isSageInParty = true
			break
		}
	}
	hasWand := false
	for _, item := range items {
		if item.name == "まほうのつえ" {
			hasWand = true
			break
		}
	}
	if !isSageInParty && !hasWand {
		return "魔王の部屋を発見できない"
	}

	hasKey := false
	for _, item := range items {
		if item.name == "くらやみのかぎ" {
			hasKey = true
			break
		}
	}
	if !hasKey {
		return "魔王の部屋を探し出す"
	}

	hasSword := false
	for _, item := range items {
		if item.name == "ひかりのつるぎ" {
			hasSword = true
			break
		}
	}
	if !hasSword {
		return "魔王の部屋に入る"
	}

	return "魔王を倒す"
}
