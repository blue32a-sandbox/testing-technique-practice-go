package ques04_test

import (
	"testing"
	"ttp/cmd/part4/ques04"
)

var dataProvider = []struct {
	job         ques04.Jobs
	gender      ques04.Gender
	personality ques04.Personality
	expected    int
}{
	// 期待値は仮
	{ques04.WarriorJob, ques04.Man, ques04.Shy, 3},
	{ques04.WizardJob, ques04.Woman, ques04.Shy, 3},
	{ques04.MonkJob, ques04.Man, ques04.Clingy, 3},
	{ques04.MonkJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.WisePerson, ques04.Woman, ques04.Clingy, 3},
	{ques04.DebaucheeJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.WarriorJob, ques04.Woman, ques04.Clingy, 3},
	{ques04.MerchantJob, ques04.Woman, ques04.LoneWolf, 3},
	{ques04.WarriorJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.FighterJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.ThiefJob, ques04.Man, ques04.LoneWolf, 3},
	{ques04.MerchantJob, ques04.Man, ques04.Clingy, 3},
	{ques04.DebaucheeJob, ques04.Woman, ques04.Careless, 3},
	{ques04.WisePerson, ques04.Woman, ques04.Tomboy, 3},
	{ques04.WizardJob, ques04.Man, ques04.Careless, 3},
	{ques04.MonkJob, ques04.Woman, ques04.Careless, 3},
	{ques04.WisePerson, ques04.Man, ques04.Careless, 3},
	{ques04.MonkJob, ques04.Woman, ques04.Tomboy, 3},
	{ques04.FighterJob, ques04.Woman, ques04.Shy, 3},
	{ques04.MerchantJob, ques04.Woman, ques04.Tomboy, 3},
	{ques04.MerchantJob, ques04.Woman, ques04.Shy, 3},
	{ques04.WarriorJob, ques04.Woman, ques04.Tomboy, 3},
	{ques04.ThiefJob, ques04.Woman, ques04.Careless, 3},
	{ques04.WisePerson, ques04.Woman, ques04.LoneWolf, 3},
	{ques04.WarriorJob, ques04.Woman, ques04.Careless, 3},
	{ques04.MerchantJob, ques04.Man, ques04.Careless, 3},
	{ques04.DebaucheeJob, ques04.Woman, ques04.Tomboy, 3},
	{ques04.FighterJob, ques04.Woman, ques04.Clingy, 3},
	{ques04.ThiefJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.ThiefJob, ques04.Woman, ques04.Tomboy, 3},
	{ques04.WizardJob, ques04.Man, ques04.Clingy, 3},
	{ques04.DebaucheeJob, ques04.Woman, ques04.Clingy, 3},
	{ques04.MonkJob, ques04.Woman, ques04.Shy, 3},
	{ques04.FighterJob, ques04.Woman, ques04.LoneWolf, 3},
	{ques04.MerchantJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.ThiefJob, ques04.Man, ques04.Clingy, 3},
	{ques04.WizardJob, ques04.Man, ques04.SlyLecher, 3},
	{ques04.DebaucheeJob, ques04.Woman, ques04.LoneWolf, 3},
	{ques04.WisePerson, ques04.Man, ques04.Shy, 3},
	{ques04.WisePerson, ques04.Man, ques04.SlyLecher, 3},
	{ques04.WarriorJob, ques04.Woman, ques04.LoneWolf, 3},
	{ques04.FighterJob, ques04.Woman, ques04.Tomboy, 3},
	{ques04.FighterJob, ques04.Man, ques04.Careless, 3},
	{ques04.WizardJob, ques04.Woman, ques04.Tomboy, 3},
}

func TestEvaluation(t *testing.T) {
	for _, data := range dataProvider {
		result := ques04.Evaluation(data.job, data.gender, data.personality)

		if result != data.expected {
			t.Fatalf("%s, %s, %sの評価が %d ではない。実際: %d",
				data.job,
				data.gender,
				data.personality,
				data.expected,
				result)
		}
	}

}
