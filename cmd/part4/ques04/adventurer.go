package ques04

type Jobs string

const (
	WarriorJob   Jobs = "戦士"
	MonkJob      Jobs = "僧侶"
	WizardJob    Jobs = "魔法使い"
	FighterJob   Jobs = "武闘家"
	MerchantJob  Jobs = "商人"
	DebaucheeJob Jobs = "遊び人"
	WisePerson   Jobs = "賢者"
	ThiefJob     Jobs = "盗賊"
)

type Gender string

const (
	Man   Gender = "男性"
	Woman Gender = "女性"
)

type Personality string

const (
	Clingy    Personality = "あまえんぼう"
	LoneWolf  Personality = "一匹狼"
	Careless  Personality = "うっかりもの"
	Tomboy    Personality = "おてんば"
	SlyLecher Personality = "むっつりスケベ"
	Shy       Personality = "引っ込み思案"
)

func Evaluation(_job Jobs, _gender Gender, _personality Personality) int {
	// TODO: 職業、性別、性格から評価値を返す
	return 3
}
