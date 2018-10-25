package main

import (
	"fmt"
	"math/rand"
	"os"

	discord "github.com/bwmarrin/discordgo"
)

func CostumeCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Here's a costume idea!")
	switch rand.Intn(4) {
	case 0:
		image("ghost.png", m.ChannelID)
		break
	case 1:
		image("skeleton.png", m.ChannelID)
		break
	case 2:
		image("witch.png", m.ChannelID)
		break
	case 3:
		image("zombie.png", m.ChannelID)
		break
	}
}

func image(image string, channelID string) {
	f, err := os.Open(image)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	ms := &discord.MessageSend{
		Embed: &discord.MessageEmbed{
			Image: &discord.MessageEmbedImage{
				URL: "attachment://" + image,
			},
		},
		Files: []*discord.File{
			&discord.File{
				Name:   image,
				Reader: f,
			},
		},
	}
	session.ChannelMessageSendComplex(channelID, ms)
}
