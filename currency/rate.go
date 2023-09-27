package currency

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/adrg/exrates"
)

func Convert(myCurr, needCurr, amountStr string) (string, error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return "", errors.New("error while parsing the amount value")
	}
	myCurr = strings.ToUpper(myCurr)
	needCurr = strings.ToUpper(needCurr)
	rates, err := exrates.Latest(needCurr, nil)
	if err != nil {
		return "", errors.New("wrong currency name")
	}
	for curr, value := range rates.Values {
		if curr == myCurr {
			return fmt.Sprintf("%.2f %s to %s is %.2f", amount, myCurr, needCurr, amount/value), nil
		}
	}
	return "", errors.New("wrong currency name")
}
