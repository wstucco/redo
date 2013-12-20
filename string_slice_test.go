package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"

	"testing"
)

func TestEach(t *testing.T) {

	Convey("Testing each", t, func() {

		slice := []string{"a", "b", "ab"}
		iterations := 0

		Each(func(s string) {
			iterations = iterations + 1
		}, slice)

		Convey(`each should be called len(slice) times`, func() {
			So(iterations, ShouldEqual, len(slice))
		})

	})

}

func TestFilter(t *testing.T) {
	Convey("Testing Filter", t, func() {

		slice := []string{"a", "A", "b", "B", "c", "C"}
		expectedResult := []string{"A", "B", "C"}

		Convey(`Should return all uppercase letters`, func() {
			result := Filter(func(s string) bool {
				return s == strings.ToUpper(s)
			}, slice)

			So(result, ShouldNotEqual, expectedResult)
			So(result, ShouldResemble, expectedResult)
		})

		Convey(`empty slice return empty slice `, func() {
			result := Filter(func(s string) bool {
				return s == strings.ToUpper(s)
			}, []string{})

			So(result, ShouldNotEqual, expectedResult)
			So(result, ShouldResemble, []string{})
		})

	})
}

func TestFirst(t *testing.T) {
	Convey("Testing StringSlice.First", t, func() {

		slice := []string{"a", "b", "ab"}
		expectedResult := "b"

		result := First(func(s string) bool {
			return s == "b"
		}, slice)

		Convey(`Should return first element equal to "b"`, func() {
			So(result, ShouldEqual, expectedResult)
		})
	})
}

func TestHead(t *testing.T) {
	Convey("Testing Head", t, func() {
		slice := []string{"a", "b", "ab"}
		expectedResult := "a"
		result := Head(slice)

		Convey(`Should return "a" from slice`, func() {
			So(result, ShouldEqual, expectedResult)
		})

		Convey(`panic if try to extract head from empty slice`, func() {
			panics := func() {
				Head(nil)
			}
			So(panics, ShouldPanicWith, "index out of range")
		})

	})
}

func TestMap(t *testing.T) {
	Convey("Testing Map", t, func() {
		slice := []string{"a", "b", "ab"}
		expectedResult := []string{"xa", "xb", "xab"}

		result := Map(func(s string) string {
			return ("x" + s)
		}, slice)

		Convey(`Map should return a list of same length as original slice`, func() {
			So(len(result), ShouldEqual, len(expectedResult))
		})

		Convey(`Map should return a new list with all elements transformed by func`, func() {
			So(result, ShouldResemble, expectedResult)
			So(result, ShouldNotEqual, expectedResult)
		})
	})
}

func TestTail(t *testing.T) {
	Convey("Testing tail", t, func() {
		slice := []string{"a", "b", "ab"}
		expectedResult := []string{"b", "ab"}

		result := Tail(slice)

		Convey(`Should return a new slice ["b", "ab"] from slice`, func() {
			So(result, ShouldNotEqual, expectedResult)
			So(result, ShouldResemble, expectedResult)
		})

		Convey(`Panic if try to extract Tail from empty slice`, func() {
			panics := func() {
				Tail(nil)
			}
			So(panics, ShouldPanicWith, "slice bounds out of range")
		})

	})
}
