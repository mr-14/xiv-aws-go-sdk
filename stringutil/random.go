package stringutil

import (
	"crypto/rand"
	"log"
	"math/big"
)

const (
	alphaBytes        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericBytes      = "0123456789"
	alphaNumericBytes = alphaBytes + numericBytes
)

func randInt(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Fatalln(err)
	}

	return n.Int64()
}

// RandomAlphaNumeric generates random alpha numberic string
func RandomAlphaNumeric(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphaNumericBytes[randInt(int64(len(alphaNumericBytes)))]
	}
	return string(b)
}

// RandomAlpha generates random string containing only alphabets
func RandomAlpha(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphaBytes[randInt(int64(len(alphaBytes)))]
	}
	return string(b)
}

// RandomNumeric generates random string containing only numbers
func RandomNumeric(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numericBytes[randInt(int64(len(numericBytes)))]
	}
	return string(b)
}
