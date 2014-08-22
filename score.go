package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Score struct {
	Joins map[string]int
	Parts map[string]int
	Quits map[string]int
}

func (s *Score) Save() {
	saveTbl(s.Joins, "joins.txt")
	saveTbl(s.Quits, "quits.txt")
	saveTbl(s.Parts, "parts.txt")
}

func saveTbl(tbl map[string]int, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer f.Close()
	for k, v := range tbl {
		if _, err = f.WriteString(strconv.Itoa(v) + "," + k + "\n"); err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
	}
}

func (s *Score) Open() {
	s.Joins = openTbl("joins.txt")
	s.Quits = openTbl("quits.txt")
	s.Parts = openTbl("parts.txt")
}

func openTbl(filename string) map[string]int {
	result := make(map[string]int)

	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, ",")
		if len(s) == 2 {
			if s[0] != "" && s[1] != "" {
				result[s[1]], _ = strconv.Atoi(s[0])
			}
		}
	}

	return result
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

	s.Open()
}

func (s *Score) AddJoin(user string) {
	s.Joins[user]++
	s.Save()
}

func (s *Score) AddPart(user string) {
	s.Parts[user]++
	s.Save()
}

func (s *Score) AddQuit(user string) {
	s.Quits[user]++
	s.Save()
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
		i++
	}

	sort.Sort(p)
	sort.Reverse(p)
	return p
}
