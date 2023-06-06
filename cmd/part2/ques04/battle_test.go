package ques04_test

import (
	"testing"
	"ttp/cmd/part2/ques04"
)

func TestCannotFindRoom(t *testing.T) {
	party := factoryPartyWithNotSage()
	items := factoryItemsWithNotKeyItem()

	result := ques04.GetProgress(party, items)

	if result != "魔王の部屋を発見できない" {
		t.Errorf("進行結果が 魔王の部屋を発見できない ではない。 actual: %s", result)
	}
}

var findRoomDataProvider = []struct {
	party ques04.Party
	items ques04.Items
}{
	{factoryPartyWithSage(), factoryItemsWithWand()},
	{factoryPartyWithSage(), factoryItemsWithNotKeyItem()},
	{factoryPartyWithNotSage(), factoryItemsWithWand()},
}

func TestFindRoom(t *testing.T) {
	for _, data := range findRoomDataProvider {
		result := ques04.GetProgress(data.party, data.items)

		if result != "魔王の部屋を探し出す" {
			t.Errorf("進行結果が 魔王の部屋を探し出す ではない。 actual: %s", result)
		}
	}
}

var enteringRoomDataProvider = []struct {
	party ques04.Party
	items ques04.Items
}{
	{factoryPartyWithSage(), factoryItemsWithWandAndKey()},
	{factoryPartyWithSage(), factoryItemsWithKey()},
	{factoryPartyWithNotSage(), factoryItemsWithWandAndKey()},
}

func TestEnteringRoom(t *testing.T) {
	for _, data := range enteringRoomDataProvider {
		result := ques04.GetProgress(data.party, data.items)

		if result != "魔王の部屋に入る" {
			t.Errorf("進行結果が 魔王の部屋に入る ではない。 actual: %s", result)
		}
	}
}

var defeatBossDataProvider = []struct {
	party ques04.Party
	items ques04.Items
}{
	{factoryPartyWithSage(), factoryItemsWithWandAndKeyAndSword()},
	{factoryPartyWithSage(), factoryItemsWithKeyAndSword()},
	{factoryPartyWithNotSage(), factoryItemsWithWandAndKeyAndSword()},
}

func TestDefeatBoss(t *testing.T) {
	for _, data := range defeatBossDataProvider {
		result := ques04.GetProgress(data.party, data.items)

		if result != "魔王を倒す" {
			t.Errorf("進行結果が 魔王を倒す ではない。 actual: %s", result)
		}
	}
}

func factoryPartyWithSage() ques04.Party {
	return ques04.Party{
		ques04.NewPartyMember("勇者"),
		ques04.NewPartyMember("戦士"),
		ques04.NewPartyMember("魔法使い"),
		ques04.NewPartyMember("賢者"),
	}
}

func factoryPartyWithNotSage() ques04.Party {
	return ques04.Party{
		ques04.NewPartyMember("勇者"),
		ques04.NewPartyMember("戦士"),
		ques04.NewPartyMember("魔法使い"),
	}
}

func factoryItemsWithNotKeyItem() ques04.Items {
	return ques04.Items{
		ques04.NewItem("やくそう"),
		ques04.NewItem("せいすい"),
	}
}

func factoryItemsWithWand() ques04.Items {
	return ques04.Items{
		ques04.NewItem("やくそう"),
		ques04.NewItem("せいすい"),
		ques04.NewItem("まほうのつえ"),
	}
}

func factoryItemsWithKey() ques04.Items {
	return ques04.Items{
		ques04.NewItem("やくそう"),
		ques04.NewItem("せいすい"),
		ques04.NewItem("くらやみのかぎ"),
	}
}

func factoryItemsWithWandAndKey() ques04.Items {
	return ques04.Items{
		ques04.NewItem("やくそう"),
		ques04.NewItem("せいすい"),
		ques04.NewItem("まほうのつえ"),
		ques04.NewItem("くらやみのかぎ"),
	}
}

func factoryItemsWithKeyAndSword() ques04.Items {
	return ques04.Items{
		ques04.NewItem("やくそう"),
		ques04.NewItem("せいすい"),
		ques04.NewItem("くらやみのかぎ"),
		ques04.NewItem("ひかりのつるぎ"),
	}
}

func factoryItemsWithWandAndKeyAndSword() ques04.Items {
	return ques04.Items{
		ques04.NewItem("やくそう"),
		ques04.NewItem("せいすい"),
		ques04.NewItem("まほうのつえ"),
		ques04.NewItem("くらやみのかぎ"),
		ques04.NewItem("ひかりのつるぎ"),
	}
}
