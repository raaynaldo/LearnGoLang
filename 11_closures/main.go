package main

import (
	"fmt"
	"strings"
)

type tokenizer func() (token string, ok bool)

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

func split(s, sep string) tokenizer {
	tokens, last := strings.Split(s, sep), 0

	return func() (string, bool) {
		if len(tokens) == last {
			return "", false
		}

		last = last + 1

		return tokens[last-1], true
	}
}

func main() {
	// sum := adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(sum(i))
	// }

	const sentence = "The quick brown fox jumps over the lazy dog"

	iter := split(sentence, " ")

	for token, ok := iter(); ok; token, ok = iter() {
		fmt.Println(token)
	}
}
