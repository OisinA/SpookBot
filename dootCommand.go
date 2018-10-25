package main

import (
	"fmt"
	"os"

	discord "github.com/bwmarrin/discordgo"
)

func DootCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	f, err := os.Open("doot.gif")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer f.Close()
	ms := &discord.MessageSend{
		Embed: &discord.MessageEmbed{
			Image: &discord.MessageEmbedImage{
				URL: "attachment://" + "doot.gif",
			},
		},
		Files: []*discord.File{
			&discord.File{
				Name:   "doot.gif",
				Reader: f,
			},
		},
	}
	fmt.Println(fmt.Sprint(ms))
	s.ChannelMessageSendComplex(m.ChannelID, ms)
}
