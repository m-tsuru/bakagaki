package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomSubsequence(s string) string {
	var builder strings.Builder

	for _, r := range s {
		if rand.Intn(2) == 1 {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	s := "バーガーキング"

	for {
		fmt.Println(randomSubsequence(s))
	}
}
