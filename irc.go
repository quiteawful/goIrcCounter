package main

import (
	"crypto/tls"
	"fmt"
	"github.com/thoj/go-ircevent"
	"strconv"
	"strings"
)

type Irc struct {
	Con      *irc.Connection
	Network  string
	Port     int
	Channels []string
}

var score *Score

func (i *Irc) Run() {
	score = &Score{}
	score.New()

	i.Con = irc.IRC("Datenkrake", "Datenkrake")
	i.Con.VerboseCallbackHandler = false
	i.Con.UseTLS = true
	i.Con.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	i.Con.Connect(i.Network + ":" + strconv.Itoa(i.Port))

	i.Con.AddCallback("001", func(e *irc.Event) {
		i.Con.Join(i.Channels[0])
	})
	i.Con.AddCallback("PRIVMSG", parseIrcMsg)
	i.Con.AddCallback("CTCP_ACTION", parseIrcMsg)

	i.Con.AddCallback("JOIN", count)
	i.Con.AddCallback("PART", count)
	i.Con.AddCallback("QUIT", count)

	i.Con.Loop()
}

func (i *Irc) WriteToChannel(content string) {
	i.Con.Privmsg(i.Channels[0], content)
}

func parseIrcMsg(e *irc.Event) {
	//user := e.Nick
	content := e.Arguments[1]
	p := strings.Split(content, " ")

	if p[0] == "!score" {
		if len(p) == 1 { // default
			printTable("Join", score.Joins)
			return
		}

		if len(p) == 2 && strings.HasPrefix(p[1], "p") {
			printTable("Part", score.Parts)
			return
		}

		if len(p) == 2 && strings.HasPrefix(p[1], "q") {
			printTable("Quit", score.Quits)
			return
		}
	}

	// control bot
	if p[0] == "!join" && len(p) == 2 && e.Nick == "marduk" {
		ctxIrc.Con.Join(p[1])
		return
	}
	if p[0] == "!part" && len(p) == 2 && e.Nick == "marduk" {
		ctxIrc.Con.Part(p[1])
		return
	}

}

func count(e *irc.Event) {
	user := e.Nick
	code := e.Code

	user = mapNickName(user)

	if user == "Datenkrake" || user == "Counter" {
		return
	}
	switch {
	case code == "JOIN":
		score.AddJoin(user)
		fmt.Println(user + " joined.")
	case code == "PART":
		score.AddPart(user)
		fmt.Println(user + " parted.")
	case code == "QUIT":
		score.AddQuit(user)
		fmt.Println(user + " quit.")
	}
}

func printTable(name string, tbl map[string]int) {
	ctxIrc.WriteToChannel("==" + name + "==")
	sorted := sortMapByValue(tbl)
	for _, v := range sorted {
		ctxIrc.WriteToChannel(strconv.Itoa(v.Value) + "  " + v.Key)
	}
}
