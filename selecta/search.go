package selecta

import "sort"

// search is best thought of as a "view" into the list of choices
// a new search object is generated on every input. it may be optimzable in
// the future

type Search struct {
	choices        []string
	matches        Matches
	index          int
	query          string
	done           bool
	visibleChoices int
}

// Match type & Matches sortable slice
type Match struct {
	value string
	score float64
}

type Matches []*Match

func (s Matches) Len() int      { return len(s) }
func (s Matches) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByScore sort method
type ByScore struct{ Matches }

func (s ByScore) Less(i, j int) bool {
	return s.Matches[i].score < s.Matches[j].score
}

func NewSearch(choices []string, index int, query string, done bool, visibleChoices int) *Search {
	search := Search{choices, make(Matches, 0, len(choices)), index, query, done, visibleChoices}

	search.createMatches()

	return &search
}

func BlankSearch(choices []string, query string, visibleChoices int) *Search {
	return NewSearch(choices, 0, query, false, 5)
}

func (s *Search) SelectedChoice() string {
	return s.choices[s.index]
}

// Create the list of matches on a Search
func (s *Search) createMatches() {
	s.matches = make(Matches, len(s.choices))

	// pair choice/score matches
	for i, choice := range s.choices {
		s.matches[i] = &Match{choice, Score(choice, s.query)}
	}

	// filter matches
	for i := len(s.matches) - 1; i >= 0; i-- {
		if s.matches[i].score == 0.0 {
			s.matches = append(s.matches[:i], s.matches[i+1:]...)
		}
	}

	// sort matches
	sort.Sort(ByScore{s.matches})
}
