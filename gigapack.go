package gigapack

import "rsc.io/quote/v3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
func Add(first, second int) int {
	return first + second
}
