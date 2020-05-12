package main

import (
	"strconv"
	"strings"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func itoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	return strings.ToUpper(strconv.FormatInt(int64(value), base))
}

func main() {
	for i := 0; i < 30; i++ {
		value := lib.RandIntBetween(-1000000, 1000000)
		base := lib.RandIntBetween(2, 16)
		lib.Challenge("ItoaBase", student.ItoaBase, itoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := lib.RandIntBetween(2, 16)
		lib.Challenge("ItoaBase", student.ItoaBase, correct.ItoaBase, lib.MaxInt, base)
		lib.Challenge("ItoaBase", student.ItoaBase, correct.ItoaBase, lib.MinInt, base)
	}
}