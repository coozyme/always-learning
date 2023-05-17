package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func GeneratePromoCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomString(n int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func main() {
	log.Println(GeneratePromoCode(7))
	log.Println(randomString(7))
}
