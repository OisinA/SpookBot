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
	var ms *discord.MessageSend
	switch rand.Intn(4) {
	case 0:
		ms = image("ghost.png")
	case 1:
		ms = image("skeleton.png")
	case 2:
		ms = image("witch.png")
	case 3:
		ms = image("zombie.png")
	}
	s.ChannelMessageSend(m.ChannelID, "Here's a costume idea!")
	s.ChannelMessageSendComplex(m.ChannelID, ms)
}

func image(image string) *discord.MessageSend {
	f, err := os.Open(image)
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	defer f.Close()
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
	return ms
}
