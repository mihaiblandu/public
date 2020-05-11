package main

import (
	"../common"
	"./student"
)

func isPositive(i int) bool {
	return i > 0
}

func isNegative0(i int) bool {
	return i < 0
}

func _map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))

	for i, el := range a {
		result[i] = f(el)
	}

	return result
}

func main() {
	functions := []func(int) bool{isPositive, isNegative0, common.IsPrime}

	type node struct {
		f   func(int) bool
		arr []int
	}

	table := []node{}

	for i := 0; i < 15; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		val := node{
			f:   function,
			arr: lib.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)
	}

	table = append(table, node{
		f:   common.IsPrime,
		arr: []int{1, 2, 3, 4, 5, 6},
	})

	for _, arg := range table {
		lib.Challenge("Map", student.Map, _map, arg.f, arg.arr)
	}
}