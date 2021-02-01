package hello

import (
	"rsc.io/quote/v3"
)

// Hello gives some greeting string
func Hello() string {
	return quote.HelloV3()
}

// Proverb gives some wisdomful string
func Proverb() string {
	return quote.Concurrency()
}
