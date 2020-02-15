package main

import (
	"strconv"
)

func main() {
}

func romanNumeral(num int) string {
	s := strconv.Itoa(num)
	bytes := ([]byte)(s)

	orderOfMagnitude := 0
	numerals := []byte{}

	for i := len(bytes) - 1; i >= 0; i-- {
		o := ordersOfMagnitude[orderOfMagnitude]
		o.value = bytes[i]
		numerals = append(o.getNumeral(), numerals...)
		orderOfMagnitude++
	}
	return string(numerals)
}

type order struct {
	i, v, x byte
	value   byte
}

func (o *order) getNumeral() []byte {
	return numeralConverterMap[o.value](o)
}

type converter func(order *order) []byte

var numeralConverterMap = map[byte]converter{
	([]byte)("0")[0]: func(order *order) []byte { return []byte{} },
	([]byte)("1")[0]: func(order *order) []byte { return []byte{order.i} },
	([]byte)("2")[0]: func(order *order) []byte { return []byte{order.i, order.i} },
	([]byte)("3")[0]: func(order *order) []byte { return []byte{order.i, order.i, order.i} },
	([]byte)("4")[0]: func(order *order) []byte { return []byte{order.i, order.v} },
	([]byte)("5")[0]: func(order *order) []byte { return []byte{order.v} },
	([]byte)("6")[0]: func(order *order) []byte { return []byte{order.v, order.i} },
	([]byte)("7")[0]: func(order *order) []byte { return []byte{order.v, order.i, order.i} },
	([]byte)("8")[0]: func(order *order) []byte { return []byte{order.v, order.i, order.i, order.i} },
	([]byte)("9")[0]: func(order *order) []byte { return []byte{order.i, order.x} },
}

var ordersOfMagnitude = []order{
	order{i: ([]byte)("I")[0], v: ([]byte)("V")[0], x: ([]byte)("X")[0]},
	order{i: ([]byte)("X")[0], v: ([]byte)("L")[0], x: ([]byte)("C")[0]},
	order{i: ([]byte)("C")[0], v: ([]byte)("D")[0], x: ([]byte)("M")[0]},
}
