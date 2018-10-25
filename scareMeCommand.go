package main

import discord "github.com/bwmarrin/discordgo"

func ScareMeCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}

	s.ChannelMessageSend(m.ChannelID, "```"+`      .-----.
    .' -   - '.
   /  .-. .-.  \    Y O U  H A V E
   |  | | | |  |    B E E N  S P O O K E D
    \ \o/ \o/ /
   _/    ^    \_
  | \  '---'  / |
  / /'--. .--'\ \
 / /'---' '---'\ \
 '.__.       .__.'
     '|     |'
      |     \
      \      '--.
       '.        '\
         '---.   |
            ,__) /`+"```")
}
