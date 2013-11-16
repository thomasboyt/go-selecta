package selecta

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func makeChoices() []string {
	f, _ := os.Open("words.txt")

	bf := bufio.NewReader(f)

	choices := make([]string, 0, 50000)
	for {
		line, _, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		word := string(line[1 : len(line)-2])
		choices = append(choices, word)
	}

	return choices
}

func generateChoices(chars string, reps int, n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < reps; j++ {
			str += chars
		}
		strs[i] = str
	}
	return strs
}

func BenchmarkNonMatching(b *testing.B) {
	choices := generateChoices("x", 16, 1000)
	query := "yyyyyyyyyyyyyyyy"

	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, query)
		}
	}
}

var query string = "xxxxxxxxxxxxxxxx"

func BenchmarkMatchingExactly(b *testing.B) {
	choices := generateChoices("x", 16, 1000)

	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, query)
		}
	}
}

func BenchmarkMatchingBrokenUp(b *testing.B) {
	choices := generateChoices("xy", 16, 1000)

	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, query)
		}
	}
}

func BenchmarkOverlappingMatches(b *testing.B) {
	choices := generateChoices("x", 40, 1000)

	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, query)
		}
	}
}

var choices []string = makeChoices()

func BenchmarkUnmatchedScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, query)
		}
	}
}

func BenchmarkMatchedScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, "ungovernableness")
		}
	}
}
