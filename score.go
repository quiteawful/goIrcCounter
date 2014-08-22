package main

import "sort"
import "fmt"

type Score struct {
	Joins map[string]int
	Parts map[string]int
	Quits map[string]int
}

func (s *Score) New() {
	s.Joins = make(map[string]int)
	s.Parts = make(map[string]int)
	s.Quits = make(map[string]int)

	nicks = make([]NickMapping, 30)

	nicks = append(nicks, NickMapping{User: "Ignite", Alias: []string{"Ignite", "Iggy", "Ig0r"}})
	nicks = append(nicks, NickMapping{User: "mayewski", Alias: []string{"mayewski", "arabayewski", "mayewski_", "mayewski__"}})
	nicks = append(nicks, NickMapping{User: "doclol", Alias: []string{"pervlol", "doclol_"}})
	nicks = append(nicks, NickMapping{User: "aimless", Alias: []string{"aimless", "aimles", "clientless"}})
	nicks = append(nicks, NickMapping{User: "nut", Alias: []string{"nut", "nutter"}})
	nicks = append(nicks, NickMapping{User: "Pitt", Alias: []string{"Pitt", "Pitt_AFK"}})
	nicks = append(nicks, NickMapping{User: "svbito", Alias: []string{"svbito", "svb", "subito"}})
	nicks = append(nicks, NickMapping{User: "Moob", Alias: []string{"Moob", "Moob_"}})
	nicks = append(nicks, NickMapping{User: "Moehre", Alias: []string{"Moehre", "Karottenkostuem"}})

	// predefined score
	s.Joins["DrHouse"] = 19
	s.Joins["Ignite"] = 10
	s.Joins["Moehre"] = 5
	s.Joins["g0bot"] = 6
	s.Joins["t-zwck"] = 2
	s.Joins["Liaf"] = 1
	s.Joins["chr0me"] = 6
	s.Joins["mayewski"] = 8
	s.Joins["Datenkatze"] = 6
	s.Joins["tabstop"] = 2
	s.Joins["aimless"] = 3
	s.Joins["svbito"] = 4
	s.Joins["Moob"] = 2
	s.Joins["MaRv"] = 1
	s.Joins["nervsack"] = 1
	s.Joins["cl1ent"] = 1
	s.Joins["nut"] = 1
	s.Joins["Pitt"] = 3

	s.Quits["DrHouse"] = 20
	s.Quits["Moehre"] = 6
	s.Quits["g0bot"] = 6
	s.Quits["Ignite"] = 8
	s.Quits["t-zwck"] = 3
	s.Quits["aimless"] = 4
	s.Quits["Liaf"] = 1
	s.Quits["chr0me"] = 6
	s.Quits["mayewski"] = 8
	s.Quits["Datenkatze"] = 6
	s.Quits["svbito"] = 4
	s.Quits["Moob"] = 2
	s.Quits["MaRv"] = 1
	s.Quits["cl1ent"] = 1
	s.Quits["nervsack"] = 1
	s.Quits["Pitt"] = 3

	s.Parts["tabstop"] = 2
	s.Parts["nut"] = 1

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

type NickMapping struct {
	User  string
	Alias []string
}

var nicks []NickMapping

func mapNickName(user string) string {
	// alle User durchgehen und mapping finden
	for _, n := range nicks {
		// alle aliase für den User rangen
		for _, alias := range n.Alias {
			if alias == user {
				return n.User // gefunden
			}
		}
	}
	// nix gefunden
	return user
}

// sort foo
type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{Key: k, Value: v}
		fmt.Printf("%q\t%q\n", m, p[i])
		i++
	}

	sort.Sort(p)
	return p
}
