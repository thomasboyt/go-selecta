# go-selecta

This is a port of @garybernhardt's [selecta](https://github.com/garybernhardt/selecta) to go. I'm making it for a few reasons:

* I needed a Go project
* Gary's been talking a lot on Twitter about optimizing Selecta, so I was curious what kind of performance boost it'd get as a compiled program instead of a Ruby script
* Selecta's code is beautiful, idiomatic Ruby, and I'm honestly curious just how ugly it gets in the translation process :)

For testing, it uses [goconvey](https://github.com/smartystreets/goconvey), which is honestly one of the best testing tools I've ever used in any language. You can just run it with `go test` or through the browser view (see the link). It also has some benchmarks for the scoring algorithim, based on the benchmark included with the original Selecta.

You'll need to run this to get convey:

```
go get github.com/smartystreets/goconvey/convey
```

### What's Done

* Scoring
* Part of searching
* Basic UI
* Input a word, backspace

### What's Not

* Actual choice input (currently uses "one, two, three")
* Scrolling
* Extra keybindings
* Lots of other cases
