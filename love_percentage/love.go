package lovepercentage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/losevs/TgBot/token"
)

type Response struct {
	Fname      string `json:"fname"`
	Sname      string `json:"sname"`
	Percentage string `json:"percentage"`
	Result     string `json:"result"`
}

func LoveCh(name1, name2 string) (string, string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://love-calculator.p.rapidapi.com/getPercentage?sname=%s&fname=%s", name1, name2), nil)
	if err != nil {
		return "", "", fmt.Errorf("API request error")
	}
	req.Header.Add("X-RapidAPI-Key", token.LOVE_API)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("client error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error while reading the body response")
	}
	Val := Response{}
	err = json.Unmarshal(body, &Val)
	if err != nil {
		return "", "", fmt.Errorf("unmarshal err: %s", err.Error())
	}

	return Val.Percentage, Val.Result, nil
}
