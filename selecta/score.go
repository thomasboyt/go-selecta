package selecta

import (
	"strings"
)

// Returns a score between 0.0 and 1.0 for a choice and a query.
func Score(choice, query string) float64 {
	if choice == "" {
		return 0.0
	}
	if query == "" {
		return 1.0
	}

	choice = strings.ToLower(choice)
	query = strings.ToLower(query)

	if len(query) > len(choice) {
		return 0.0
	}

	length := computeMatchLength(choice, query)

	if length == -1 {
		return 0.0
	}

	score := float64(len(query)) / float64(length)
	return score / float64(len(choice))
}

// Computes the length of the shortest substring between the first char and
// the rest of the characters. Returns -1 if the query isn't within the string
func computeMatchLength(str, charsStr string) int {
	maxLength := -1

	chars := []rune(charsStr)
	firstChar := chars[0]
	restChars := chars[1:]

	firstIndexes := findIndexesInString(str, firstChar)
	for _, firstIndex := range firstIndexes {
		lastIndex := findEndOfMatch(str, restChars, firstIndex)
		if lastIndex != -1 {
			length := lastIndex - firstIndex + 1
			if maxLength == -1 || length < maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

// Find all occurances of a char in a string
func findIndexesInString(str string, char rune) []int {
	indexes := make([]int, 0)

	for i, runeValue := range str {
		if runeValue == char {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

// Find each of the characters in the string, moving ltr.
func findEndOfMatch(str string, restChars []rune, firstIndex int) int {
	lastIndex := firstIndex
	for _, charRuneValue := range restChars {
		charIndex := -1
		for j, runeValue := range str[lastIndex:] {
			if runeValue == charRuneValue {
				charIndex = j
				break
			}
		}

		// Return -1 if a character isn't found at all
		if charIndex == -1 {
			return -1
		}

		lastIndex = lastIndex + charIndex
	}
	return lastIndex
}
