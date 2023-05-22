package ques03

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

func ConvertTatamikazuToArea(tatamikazu int64) (string, error) {
	if tatamikazu <= -129 || tatamikazu >= 128 {
		return "", errors.New("畳数は-128以上127以下を入力してください")
	}
	if tatamikazu <= 0 {
		return "畳数は１以上を入力してください", nil
	}

	dTatamikazu := decimal.NewFromInt(tatamikazu)
	area := dTatamikazu.Mul(decimal.NewFromFloat(1.65))

	return fmt.Sprintf("%s㎡", area.String()), nil
}
