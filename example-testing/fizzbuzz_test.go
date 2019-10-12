package demo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	Convey("Given integers not divisible by 3 or 5", t, func() {

		Convey("it should be return number", func() {
			So(fizzbuzz(1), ShouldEqual, "1")
			So(fizzbuzz(2), ShouldEqual, "2")
			So(fizzbuzz(4), ShouldEqual, "4")
		})

	})

	Convey("For multiples of 3", t, func() {

		Convey("it should  return \"fizz\"", func() {
			So(fizzbuzz(3), ShouldEqual, "fizz")
			So(fizzbuzz(6), ShouldEqual, "fizz")
			So(fizzbuzz(99), ShouldEqual, "fizz")
		})

	})

	Convey("For multiples of 5", t, func() {

		Convey("it should return \"buzz\"", func() {
			So(fizzbuzz(5), ShouldEqual, "buzz")
			So(fizzbuzz(25), ShouldEqual, "buzz")
			So(fizzbuzz(55), ShouldEqual, "buzz")
		})

	})

	Convey("For multiples of both 3 and 5", t, func() {

		Convey("it should return \"fizzbuzz\"", func() {
			So(fizzbuzz(15), ShouldEqual, "fizzbuzz")
		})

	})

}
