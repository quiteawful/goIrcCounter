package main

type Score struct {
	Joins map[string]int
	Parts map[string]int
	Quits map[string]int
}

func (s *Score) New() {
	s.Joins = make(map[string]int)
	s.Parts = make(map[string]int)
	s.Quits = make(map[string]int)
}

func (s *Score) AddJoin(user string) {
	s.Joins[user]++
}

func (s *Score) AddPart(user string) {
	s.Parts[user]++
}

func (s *Score) AddQuit(user string) {
	s.Quits[user]++
}
