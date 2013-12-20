package main

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestReplaceBaseName(t *testing.T) {
	Convey("Testing replaceBaseName", t, func() {

		Convey(`"redo.do" should return "default.do"`, func() {
			So(replaceBaseName("redo.do", "default"), ShouldEqual, "default.do")
		})

		Convey(`"./redo.do" should return "default.do", current folder should be stripped`, func() {
			So(replaceBaseName("./redo.do", "default"), ShouldEqual, "default.do")
		})

		Convey(`"./redo.do" should not be "./default.do", current folder should be stripped`, func() {
			So(replaceBaseName("./redo.do", "./default"), ShouldNotEqual, "./default.do")
		})
	})
}

func TestHasExtension(t *testing.T) {
	Convey("Testing hasEsxtension", t, func() {

		So(hasExtension("redo.do"), ShouldBeTrue)
		So(hasExtension("redo"), ShouldBeFalse)

	})
}

func TestFileExists(t *testing.T) {
	Convey("Testing fileExists", t, func() {

		So(fileExists("redo.do"), ShouldBeTrue)
		So(fileExists("missing"), ShouldBeFalse)

		Convey("should return false on Dirs", func() {
			So(fileExists("/etc"), ShouldBeFalse)
		})
	})
}
