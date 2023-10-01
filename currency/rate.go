package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/losevs/TgBot/token"
)

func Convert(myCurr, needCurr, amountStr string) (string, error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return "", errors.New("error while parsing the amount value")
	}
	myCurr = strings.ToUpper(myCurr)
	needCurr = strings.ToUpper(needCurr)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s_%s&compact=ultra&apiKey=%s", needCurr, myCurr, token.CONVERT_API), nil)
	if err != nil {
		return "", fmt.Errorf("API request error")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	var data map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f %s to %s is %.2f\nrate = %.4f", amount, myCurr, needCurr, amount/data[fmt.Sprintf("%s_%s", needCurr, myCurr)], data[fmt.Sprintf("%s_%s", needCurr, myCurr)]), nil
}
