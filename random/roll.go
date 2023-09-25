package random

import (
	"errors"
	"math/rand"
	"strconv"
)

func Roll(number string) (int, error) {
	numInt, err := strconv.Atoi(number)
	if numInt < 1 {
		return 0, errors.New("Less than zero")
	}
	if err != nil {
		return 0, errors.New("Bad request - error while parsing the value")
	}
	randomNum := rand.Intn(numInt + 1)
	return randomNum, nil
}
