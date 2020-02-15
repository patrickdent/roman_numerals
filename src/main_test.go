package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIt(t *testing.T) {
	Convey("It converts to Roman numerals", t, func() {
		Convey("converts 1", func() {
			numeral := romanNumeral(1)
			So(numeral, ShouldEqual, "I")
		})
		Convey("converts 3", func() {
			numeral := romanNumeral(3)
			So(numeral, ShouldEqual, "III")
		})
		Convey("converts 4", func() {
			numeral := romanNumeral(4)
			So(numeral, ShouldEqual, "IV")
		})
		Convey("converts 8", func() {
			numeral := romanNumeral(8)
			So(numeral, ShouldEqual, "VIII")
		})
		Convey("converts 9", func() {
			numeral := romanNumeral(9)
			So(numeral, ShouldEqual, "IX")
		})
		Convey("converts 10", func() {
			numeral := romanNumeral(10)
			So(numeral, ShouldEqual, "X")
		})
		Convey("converts 84", func() {
			numeral := romanNumeral(84)
			So(numeral, ShouldEqual, "LXXXIV")
		})
		Convey("converts 98", func() {
			numeral := romanNumeral(98)
			So(numeral, ShouldEqual, "XCVIII")
		})
	})
}
