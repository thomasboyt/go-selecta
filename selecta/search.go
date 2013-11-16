package selecta

import (
	"regexp"
	"sort"
)

// search is best thought of as a "view" into the list of choices
// a new search object is generated on every input. it may be optimzable in
// the future

type Search struct {
	choices        []string
	Matches        Matches
	Index          int
	Query          string
	Done           bool
	visibleChoices int
}

// Match type & Matches sortable slice
type Match struct {
	Value string
	score float64
}

type Matches []*Match

func (s Matches) Len() int      { return len(s) }
func (s Matches) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByScore sort method
type ByScore struct{ Matches }

func (s ByScore) Less(i, j int) bool {
	return s.Matches[i].score > s.Matches[j].score
}

// Search constructors
func NewSearch(choices []string, index int, query string, done bool, visibleChoices int) *Search {
	search := Search{choices, make(Matches, 0, len(choices)), index, query, done, visibleChoices}

	search.createMatches()

	return &search
}

func BlankSearch(choices []string, query string, visibleChoices int) *Search {
	return NewSearch(choices, 0, query, false, visibleChoices)
}

// Search methods
func (s *Search) SelectedChoice() string {
	if len(s.Matches) == 0 {
		return ""
	}
	return s.Matches[s.Index].Value
}

func (s *Search) AppendQuery(str string) {
	s.Query += str
	s.Index = 0
	s.createMatches()
}

func (s *Search) Backspace() {
	if len(s.Query) == 0 {
		return
	}
	s.Query = s.Query[:len(s.Query)-1]
	s.Index = 0
	s.createMatches()
}

var wordRegexp *regexp.Regexp = regexp.MustCompile("[^ ]* *$")

func (s *Search) DeleteWord() {
	s.Query = wordRegexp.ReplaceAllString(s.Query, "")
}

func (s *Search) Down() {
	nextIndex := s.Index + 1
	if nextIndex == len(s.Matches) {
		return
	} else if nextIndex == s.visibleChoices {
		return
	}

	s.Index = nextIndex
}

func (s *Search) Up() {
	nextIndex := s.Index - 1
	if nextIndex < 0 {
		return
	}

	s.Index = nextIndex
}

// Create the list of matches on a Search
func (s *Search) createMatches() {
	s.Matches = make(Matches, len(s.choices))

	// pair choice/score Matches
	for i, choice := range s.choices {
		s.Matches[i] = &Match{choice, Score(choice, s.Query)}
	}

	// filter Matches
	for i := len(s.Matches) - 1; i >= 0; i-- {
		if s.Matches[i].score == 0.0 {
			s.Matches = append(s.Matches[:i], s.Matches[i+1:]...)
		}
	}

	// sort Matches
	sort.Sort(ByScore{s.Matches})
}
