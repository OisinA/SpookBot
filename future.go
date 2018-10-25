package main

import (
	"math/rand"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

func FutureCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	if rand.Intn(5) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Your future looks spooookyyyy!")
		go SpookyFuture(s, m.Author)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Your future looks booooorrrrrinnnnggggg")
		s.ChannelMessageSend(m.ChannelID, "Someone else's future is spooky however")
		user := getRandomUser(s, m).User
		go SpookyFuture(s, user)
	}

}

func SpookyFuture(s *discord.Session, user *discord.User) {
	random := rand.Int63n(1800)
	time.Sleep(time.Duration(random * time.Second.Nanoseconds()))
	ch, err := s.UserChannelCreate(user.ID)
	if err != nil {
		return
	}
	s.ChannelMessageSend(ch.ID, "BOO!")
}
