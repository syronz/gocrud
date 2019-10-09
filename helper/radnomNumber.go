package helper

import (
	"math/rand"
	"strconv"
	"time"
)

// RandomNumber is used for generating numbers for replacing inside payloads.
func RandomNumber(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min+1) + min
	strNum := strconv.Itoa(num)

	return strNum
}
