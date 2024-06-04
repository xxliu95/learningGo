package Random

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

func RandRunes(n int, runes string) string {
	rand.Seed(time.Now().UnixNano())

	nRunes := utf8.RuneCountInString(runes)
	var result string = ""
	for i := 0; i < n; i++ {
		randNum := int(rand.Float64() * float64(nRunes))
		result += string([]rune(runes)[randNum])
	}
	return result
}
