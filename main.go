package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
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

	err = session.UpdateStatus(0, "*aaaaaahhhh*")
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(ParseCommands)

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
