package demo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringSliceEqual(t *testing.T) {

	Convey("Given a != nil and b!=nil", t, func() {

		Convey("it should be return true", func() {
			a := []string{"hello"}
			b := []string{"hello"}
			So(stringslice(a, b), ShouldBeTrue)
		})

	})

	Convey("Given a = nil and b = nil", t, func() {

		Convey("it should be return true", func() {
			a := []string(nil)
			b := []string(nil)
			So(stringslice(a, b), ShouldBeTrue)
		})

	})

	Convey("Given a = nil and b!=nil", t, func() {

		Convey("it should be return false", func() {
			a := []string(nil)
			b := []string{}
			So(stringslice(a, b), ShouldBeFalse)
		})

	})

	Convey("Given len(a) != len(b)", t, func() {

		Convey("it should be return false", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(stringslice(a, b), ShouldBeFalse)
		})

	})

	Convey("Comparing two variables", t, func() {
		myVar := "Hello, world!"

		Convey(`"Asdf" should NOT equal "qwerty"`, func() {
			So("Asdf", ShouldNotEqual, "qwerty")
		})

		Convey("myVar should not be nil", func() {
			So(myVar, ShouldNotBeNil)
		})
		Convey("This isn't yet implemented", nil)
	})

}
