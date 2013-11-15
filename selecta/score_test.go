package selecta

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestScoring(t *testing.T) {

	Convey("Basic matching", t, func() {
		Convey("scores 0 when the choice is empty", func() {
			So(Score("", "foo"), ShouldEqual, 0.0)
		})
		Convey("scores 1 when the query is empty", func() {
			So(Score("foo", ""), ShouldEqual, 1.0)
		})
		Convey("scores 0 when the query is longer than the choice", func() {
			So(Score("foo", "foobar"), ShouldEqual, 0.0)
		})
		Convey("scores 0 when the query doesn't match at all", func() {
			So(Score("a", "b"), ShouldEqual, 0.0)
		})
		Convey("scores 0 when only a prefix of the query matches", func() {
			So(Score("ab", "ac"), ShouldEqual, 0.0)
		})
		Convey("scores greater than 0 when it matches", func() {
			So(Score("a", "a"), ShouldBeGreaterThan, 0.0)
			So(Score("ab", "a"), ShouldBeGreaterThan, 0.0)
			So(Score("ba", "a"), ShouldBeGreaterThan, 0.0)
			So(Score("bab", "a"), ShouldBeGreaterThan, 0.0)
			So(Score("babababab", "aaaa"), ShouldBeGreaterThan, 0.0)
		})
		Convey("scores 1, normalized to length, when the query equals the choice", func() {
			So(Score("a", "a"), ShouldEqual, 1.0)
			So(Score("ab", "ab"), ShouldEqual, 0.5)
			So(Score("a long string", "a long string"), ShouldEqual, 1.0 / float64(len("a long string")))
			So(Score("spec/search_spec.rb", "sear"), ShouldEqual, 1.0 / float64(len("spec/search_spec.rb")))
		})

	})

	Convey("Character matching", t, func() {
		Convey("matches punctuation", func() {
			So(Score("/! symbols $^", "/!S^"), ShouldBeGreaterThan, 0.0)
		})
		Convey("is case insensitive", func() {
			So(Score("a", "A"), ShouldEqual, 1.0)
			So(Score("A", "a"), ShouldEqual, 1.0)
		})
		Convey("doesn't match when the same letter is repeated in the choice", func() {
			So(Score("a", "aa"), ShouldEqual, 0.0)
		})
	})

	Convey("Match quality", t, func() {
		Convey("scores higher for better matches", func() {
			So(Score("selecta.gemspec", "asp"), ShouldBeGreaterThan, Score("algorithm4_spec.rb", "asp"))
			So(Score("README.md", "em"), ShouldBeGreaterThan, Score("benchmark.rb", "em"))
			So(Score("search.rb", "sear"), ShouldBeGreaterThan, Score("spec/search_spec.rb", "sear"))
		})
		Convey("scores shorter matches higher", func() {
			So(Score("fbb", "fbb"), ShouldBeGreaterThan, Score("foo bar baz", "fbb"))
			So(Score("foo", "foo"), ShouldBeGreaterThan, Score("longer foo", "foo"))
			So(Score("foo", "foo"), ShouldBeGreaterThan, Score("foo longer", "foo"))
		})
		Convey("sometimes scores longer strings higher if they have a better match", func() {
			So(Score("long 12 long", "12"), ShouldBeGreaterThan, Score("1 long 2", "12"))
		})
		Convey("scores the tighter of two matches, regardless of order", func() {
			tight := "12"
			loose := "1padding2"
			So(Score(tight + loose, "12"), ShouldEqual, 1.0 / float64(len(tight + loose)))
			So(Score(loose + tight, "12"), ShouldEqual, 1.0 / float64(len(loose + tight)))
		})
	})

}
