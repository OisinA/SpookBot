package main

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Execute     func(*discord.Session, *discord.MessageCreate, string)
}

var commands = make(map[string]Command)

func RegisterCommands() {
	commands[".doot"] = Command{"doot", "doots", DootCommand}
	commands[".spookify"] = Command{"spookify", "make your name spooky", SpookifyCommand}
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
		fmt.Println("is not ok " + command)
		return
	}

	returned.Execute(s, m, strings.Join(split[1:], " "))
	return
}
