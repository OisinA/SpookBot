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

var commands = make(map[string]Command)

func RegisterCommands() {
	commands[".doot"] = Command{"doot", "doots", DootCommand}
	commands[".spookify"] = Command{"spookify", "make your name spooky", SpookifyCommand}
	commands[".scareme"] = Command{"scareme", "be scared", ScareMeCommand}
	commands[".spook"] = Command{"spook", "spook your friends!", SpookCommand}
	commands[".future"] = Command{"future", "predict your future", FutureCommand}
	commands[".costume"] = Command{"costume", "sends a costume idea", CostumeCommand}
	commands[".help"] = Command{"help", "help your spooks", HelpCommand}
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

func HelpCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if len(message) == 0 {
		keys := make([]string, 0, len(commands))
		for k := range commands {
			keys = append(keys, k)
		}
		s.ChannelMessageSend(m.ChannelID, "To use a command, send a message as follows: .[command]\nTo find out more about a command, use .help [command].\nAvailable commands: "+strings.Join(keys, ", "))
	} else {
		c, ok := commands["."+message]
		if !ok {
			s.ChannelMessageSend(m.ChannelID, "Command not found.")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "."+message+"\nDescription: "+c.Description)
	}
}
