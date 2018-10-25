package main

import (
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Execute     func(*discord.Session, *discord.MessageCreate, string)
}

var commands = make(map[string]command)

func RegisterCommands() {

}

func ParseCommands(s *discord.Session, m *discord.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := m.Message.Content
	split := strings.Split(content, " ")
	command := strings.ToLower(split[0])

	returned, ok := commands[command]
	if !ok {
		return
	}

	returned.Execute(s, m, strings.Join(split[1:], " "))
	return
}
