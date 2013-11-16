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
		// TODO
	})

	Convey("Moving up the list", t, func() {
		// TODO
	})

	Convey("Backspaces over characters", t, func() {
		// TODO
	})

	Convey("Deletes words", t, func() {
		// TODO
	})

	Convey("Matching", t, func() {
		Convey("only returns matching choices", func() {
			search := BlankSearch(choices, "one", 0)
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
