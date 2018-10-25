package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
)

var session *discord.Session

func main() {
	var err error
	session, err := discord.New("Bot " + ReadToken())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	err = session.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	RegisterCommands()

	err = session.UpdateStatus(0, "*aaaaaahhhh*")
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(ParseCommands)
	session.AddHandler(DidIHearSpook)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc
}

func ReadToken() string {
	file := "token"
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(dat)
}

func DidIHearSpook(s *discord.Session, m *discord.MessageCreate) {
	if m.Author.Bot {
		return
	}
	content := m.Message.Content
	split := strings.Split(content, " ")
	didTheySpook := false
	for _, i := range split {
		if strings.ToLower(i) == "spook" || strings.ToLower(i) == "spooked" || strings.ToLower(i) == "spooks" {
			didTheySpook = true
		}
	}

	if didTheySpook {
		s.ChannelMessageSend(m.ChannelID, "Did I hear spook??")
	}
}
