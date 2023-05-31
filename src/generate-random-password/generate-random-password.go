package pasawutil

import (
	"errors"
	"math/rand"

	"github.com/sethvargo/go-password/password"
)

func GenerateRandomPassword(length int) (string, error) {
	if length == 0 {
		return "", errors.New("no length provided")
	}

	if length <= 10 {
		return "", errors.New("length must be greater than 10")
	}

	numberOfDigits := randomInt(1, length/randomInt(3, 7))
	numberOfSymbols := randomInt(1, length/randomInt(3, 7))

	return password.Generate(length, numberOfDigits, numberOfSymbols, false, false)
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
