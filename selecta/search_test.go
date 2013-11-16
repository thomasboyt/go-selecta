package selecta

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearching(t *testing.T) {

	choices := []string{"one", "two", "three"}

	Convey("Selects the first choice by default", t, func() {
		search := BlankSearch(choices, "", 0)
		So(search.SelectedChoice(), ShouldEqual, "one")
	})

	Convey("Moving down the list", t, func() {
		Convey("moves down the list", func() {
			search := BlankSearch(choices, "", 5)
			search.Down()
			So(search.SelectedChoice(), ShouldEqual, "two")
		})

		Convey("won't move past the end of the list", func() {
			search := BlankSearch(choices, "", 5)
			search.Down()
			search.Down()
			search.Down()
			search.Down()
			So(search.SelectedChoice(), ShouldEqual, "three")
		})

		Convey("won't move past the visible choice limit", func() {
			search := BlankSearch(choices, "", 2)
			search.Down()
			search.Down()
			search.Down()
			So(search.SelectedChoice(), ShouldEqual, "two")
		})

		Convey("moves down the filtered search results", func() {
			search := BlankSearch(choices, "", 5)
			search.AppendQuery("t")
			search.Down()
			So(search.SelectedChoice(), ShouldEqual, "three")
		})
	})

	Convey("Moving up the list", t, func() {
		search := BlankSearch(choices, "", 5)
		search.Down()
		search.Up()
		So(search.SelectedChoice(), ShouldEqual, "one")
	})

	Convey("Backspaces over characters", t, func() {
		search := BlankSearch(choices, "", 5)
		search.AppendQuery("e")
		So(search.Query, ShouldEqual, "e")
		search.Backspace()
		So(search.Query, ShouldEqual, "")
	})

	Convey("Deletes words", t, func() {
		search := BlankSearch(choices, "", 5)

		search.AppendQuery("a")
		search.DeleteWord()
		So(search.Query, ShouldEqual, "")

		search.AppendQuery("a ")
		search.DeleteWord()
		So(search.Query, ShouldEqual, "")

		search.AppendQuery("a b")
		search.DeleteWord()
		So(search.Query, ShouldEqual, "a ")

		search = BlankSearch(choices, "", 5)
		search.AppendQuery("a b ")
		search.DeleteWord()
		So(search.Query, ShouldEqual, "a ")

		search = BlankSearch(choices, "", 5)
		search.AppendQuery("a b ")
		search.DeleteWord()
		So(search.Query, ShouldEqual, "a ")

		search = BlankSearch(choices, "", 5)
		search.AppendQuery(" a b")
		search.DeleteWord()
		So(search.Query, ShouldEqual, " a ")
	})

	Convey("Matching", t, func() {
		Convey("only returns matching choices", func() {
			search := BlankSearch(choices, "", 0)
			search.AppendQuery("one")
			So(len(search.Matches), ShouldEqual, 1)
			So(search.Matches[0].Value, ShouldEqual, "one")
		})
		Convey("sorts the choices by score", func() {
			search := BlankSearch([]string{"selecta.go", "selecta_test.go"}, "", 0)
			search.AppendQuery("se")
			So(search.Matches[0].Value, ShouldEqual, "selecta.go")
			So(search.Matches[1].Value, ShouldEqual, "selecta_test.go")
		})
	})
}
