package selecta

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func getChoices() []string {
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

func BenchmarkUnmatchedScore(b *testing.B) {
	choices := getChoices()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, "xxxxxxxxxxxxxxxx")
		}
	}
}

func BenchmarkMatchedScore(b *testing.B) {
	choices := getChoices()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, choice := range choices {
			Score(choice, "ungovernableness")
		}
	}
}
