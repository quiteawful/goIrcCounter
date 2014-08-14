package main

import (
	"crypto/tls"
	"github.com/thoj/go-ircevent"
	"strconv"
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

	i.Con = irc.IRC("Counter", "Datenkrake")
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

	if content == "!score" {
		// Joins
		ctxIrc.WriteToChannel("#  -  Name")
		ctxIrc.WriteToChannel("==Joins==")
		for k, v := range score.Joins {
			ctxIrc.WriteToChannel(strconv.Itoa(v) + "  " + k)
		}
		// Parts
		ctxIrc.WriteToChannel("==Parts==")
		for k, v := range score.Parts {
			ctxIrc.WriteToChannel(strconv.Itoa(v) + "  " + k)
		}
		// Quits
		ctxIrc.WriteToChannel("==Quits==")
		for k, v := range score.Quits {
			ctxIrc.WriteToChannel(strconv.Itoa(v) + "  " + k)
		}
	}
}

func count(e *irc.Event) {
	user := e.Nick
	code := e.Code

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
