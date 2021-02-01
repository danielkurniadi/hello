package hello

import (
	"strings"

	"rsc.io/quote/v3"
)

// Hello gives some greeting string
func Hello(args ...string) string {
	if len(args) == 0 {
		return quote.HelloV3()
	}
	return quote.HelloV3() + " " + strings.Join(args, " ")
}

// Proverb gives some wisdomful string
func Proverb() string {
	return quote.Concurrency()
}
