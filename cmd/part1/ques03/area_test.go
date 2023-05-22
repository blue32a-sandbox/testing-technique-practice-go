package ques03_test

import (
	"testing"
	"ttp/cmd/part1/ques03"
)

var dataProvider = []struct {
	tatamikazu int64
	actual     string
}{
	{-128, "畳数は１以上を入力してください"},
	{-100, "畳数は１以上を入力してください"},
	{0, "畳数は１以上を入力してください"},
	{1, "1.65㎡"},
	{100, "165㎡"},
	{127, "209.55㎡"},
}

func TestConvert(t *testing.T) {
	for _, data := range dataProvider {
		result, _ := ques03.ConvertTatamikazuToArea(data.tatamikazu)

		if result != data.actual {
			t.Fatalf(
				"畳数が %d のとき表示メッセージが %s ではない。 actual: %s",
				data.tatamikazu,
				data.actual,
				result)
		}
	}
}

var errorDataProvider = []struct {
	tatamikazu int64
}{
	{-200},
	{-129},
	{128},
	{200},
}

func TestInvalidTatamikazu(t *testing.T) {
	for _, data := range errorDataProvider {
		_, err := ques03.ConvertTatamikazuToArea(data.tatamikazu)

		if err == nil {
			t.Fatalf("畳数が %d のときエラーが発生しない。", data.tatamikazu)
		}
	}
}
