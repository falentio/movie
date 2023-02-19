package random

import "math/rand"

const (
	vowel = "aiueo"
	consonant = "bcdfghjklmnpqrstvwxyz"
)

func Slug(l int) string {
	var result []byte
	i := rand.Int()
	for len(result) < l {
		chars := vowel
		if i % 2 == 0 {
			chars = consonant
		}

		c := chars[rand.Intn(len(chars))]
		result = append(result, c)
		i++
	}
	return string(result)
}