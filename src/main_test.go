package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIt(t *testing.T) {
	Convey("It converts to Roman numerals", t, func() {
		Convey("converts 1", func() {
			numeral, _ := romanNumeral(1)
			So(numeral, ShouldEqual, "I")
		})
		Convey("converts 3", func() {
			numeral, _ := romanNumeral(3)
			So(numeral, ShouldEqual, "III")
		})
		Convey("converts 4", func() {
			numeral, _ := romanNumeral(4)
			So(numeral, ShouldEqual, "IV")
		})
		Convey("converts 8", func() {
			numeral, _ := romanNumeral(8)
			So(numeral, ShouldEqual, "VIII")
		})
		Convey("converts 9", func() {
			numeral, _ := romanNumeral(9)
			So(numeral, ShouldEqual, "IX")
		})
		Convey("converts 10", func() {
			numeral, _ := romanNumeral(10)
			So(numeral, ShouldEqual, "X")
		})
		Convey("converts 84", func() {
			numeral, _ := romanNumeral(84)
			So(numeral, ShouldEqual, "LXXXIV")
		})
		Convey("converts 98", func() {
			numeral, _ := romanNumeral(98)
			So(numeral, ShouldEqual, "XCVIII")
		})
		Convey("converts 1983", func() {
			numeral, _ := romanNumeral(1983)
			So(numeral, ShouldEqual, "MCMLXXXIII")
		})
		Convey("converts 2020", func() {
			numeral, _ := romanNumeral(2020)
			So(numeral, ShouldEqual, "MMXX")
		})
		Convey("can't go above 3999", func() {
			_, err := romanNumeral(4000)
			So(err.Error(), ShouldEqual, "I can't count that high")
		})
	})
}
