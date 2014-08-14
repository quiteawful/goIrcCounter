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

	nicks = make([]NickMapping, 30)

	nicks = append(nicks, NickMapping{User: "Ignite", Alias: []string{"Ignite", "Iggy", "Ig0r"}})
	nicks = append(nicks, NickMapping{User: "mayewski", Alias: []string{"mayewski", "arabayewski", "mayewski_", "mayewski__"}})
	nicks = append(nicks, NickMapping{User: "doclol", Alias: []string{"pervlol", "doclol_"}})
	nicks = append(nicks, NickMapping{User: "aimless", Alias: []string{"aimless", "aimles"}})
	nicks = append(nicks, NickMapping{User: "nut", Alias: []string{"nut", "nutter"}})
	nicks = append(nicks, NickMapping{User: "Pitt", Alias: []string{"Pitt", "Pitt_AFK"}})
	nicks = append(nicks, NickMapping{User: "svbito", Alias: []string{"svbito", "svb", "subito"}})

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
	return "deine Mama"
}
