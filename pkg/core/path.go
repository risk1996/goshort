package core

import "math/rand"

func RandShortLinkPath(length int) string {
	// cSpell:disable-next-line
	var chars = []rune("bcdfghjklmnpqrstvwxyz0123456789")
	res := ""

	for i := 0; i < length; i++ {
		res += string(chars[rand.Intn(len(chars))])
	}

	return res
}
