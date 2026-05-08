package main

import (
	"strconv"
	"strings"
)

func Bank(greeting string) string {
	greeting = strings.ToLower(greeting)
	startsWithHello := strings.HasPrefix(greeting, "hello")
	amount := 100
	if startsWithHello {
		amount = 0
	} else if greeting[0] == 'h' && greeting != "hello" {
		amount = 20
	}
	return "$" + strconv.Itoa(amount)
}
