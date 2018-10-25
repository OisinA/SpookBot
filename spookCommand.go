package main

import (
	"math/rand"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func SpookCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}

	var user *discord.Member

	if message == "" {
		user = getRandomUser(s, m)
	} else {
		user = getUser(s, m, message)
	}

	if user == nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't find the user. Try spooking again!")
		return
	}

	ch, err := s.UserChannelCreate(user.User.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Something went wrong :(")
		return
	}
	s.ChannelMessageSend(ch.ID, "Someone is thinking of you... And they decided to spook you. SPOOK!")
	s.ChannelMessageSend(m.ChannelID, "They have been spooked <3")
}

func getUser(s *discord.Session, m *discord.MessageCreate, username string) *discord.Member {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return nil
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return nil
	}

	for _, u := range g.Members {
		if "<@"+u.User.ID+">" == strings.Replace(username, "!", "", 1) {
			return u
		}
	}
	return nil
}

func getRandomUser(s *discord.Session, m *discord.MessageCreate) *discord.Member {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return nil
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return nil
	}

	return g.Members[rand.Intn(len(g.Members))]
}
