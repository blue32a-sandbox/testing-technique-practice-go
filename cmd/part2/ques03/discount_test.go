package ques03_test

import (
	"testing"
	"ttp/cmd/part2/ques03"
)

func TestMoreThanSevenItemsWithWaiShirtAndTie(t *testing.T) {
	items := FactoryItemsMoreThanSevenWithWaiShirtAndTie()

	result := ques03.GetDiscountRate(items)

	if result != 12 {
		t.Fatalf("割引率が 12 %%ではない。 actual: %d", result)
	}
}

func TestMoreThanSevenItemsWithWaiShirtNoTie(t *testing.T) {
	items := FactoryItemsMoreThanSevenWithWaiShirtNoTie()

	result := ques03.GetDiscountRate(items)

	if result != 7 {
		t.Fatalf("割引率が 7 %%ではない。 actual: %d", result)
	}
}

func TestMoreThanSevenItemsWithNoWaiShirtNoTie(t *testing.T) {
	items := FactoryItemsMoreThanSevenWithNoWaiShirtNoTie()

	result := ques03.GetDiscountRate(items)

	if result != 7 {
		t.Fatalf("割引率が 7 %%ではない。 actual: %d", result)
	}
}

func TestLessThanSevenItemsWithWaiShirtAndTie(t *testing.T) {
	items := FactoryItemsLessThanSevenWithWaiShirtAndTie()

	result := ques03.GetDiscountRate(items)

	if result != 5 {
		t.Fatalf("割引率が 5 %%ではない。 actual: %d", result)
	}
}

func TestLessThanSevenItemsWithWaiShirtNoTie(t *testing.T) {
	items := FactoryItemsLessThanSevenWithWaiShirtNoTie()

	result := ques03.GetDiscountRate(items)

	if result != 0 {
		t.Fatalf("割引率が 0 %%ではない。 actual: %d", result)
	}
}

func TestLessThanSevenItemsWithNoWaiShirtNoTie(t *testing.T) {
	items := FactoryItemsLessThanSevenWithNoWaiShirtNoTie()

	result := ques03.GetDiscountRate(items)

	if result != 0 {
		t.Fatalf("割引率が 0 %%ではない。 actual: %d", result)
	}
}

func FactoryItemsMoreThanSevenWithWaiShirtAndTie() []ques03.Item {
	return []ques03.Item{
		ques03.NewItem(ques03.WAI_SHIRT),
		ques03.NewItem(ques03.TIE),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
	}
}

func FactoryItemsMoreThanSevenWithWaiShirtNoTie() []ques03.Item {
	return []ques03.Item{
		ques03.NewItem(ques03.WAI_SHIRT),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
	}
}

func FactoryItemsMoreThanSevenWithNoWaiShirtNoTie() []ques03.Item {
	return []ques03.Item{
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
	}
}

func FactoryItemsLessThanSevenWithWaiShirtAndTie() []ques03.Item {
	return []ques03.Item{
		ques03.NewItem(ques03.WAI_SHIRT),
		ques03.NewItem(ques03.TIE),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
	}
}

func FactoryItemsLessThanSevenWithWaiShirtNoTie() []ques03.Item {
	return []ques03.Item{
		ques03.NewItem(ques03.WAI_SHIRT),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
	}
}

func FactoryItemsLessThanSevenWithNoWaiShirtNoTie() []ques03.Item {
	return []ques03.Item{
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
		ques03.NewItem(ques03.OTHER),
	}
}
