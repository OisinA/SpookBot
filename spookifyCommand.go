package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	discord "github.com/bwmarrin/discordgo"
)

type SpookyWords struct {
	Words []string
}

var (
	spookyWords = []string{}
)

func SpookifyCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}

	if len(spookyWords) == 0 {
		spookyWords = LoadSpookyWords()
	}

	chosenWords := []string{}

	for _, i := range spookyWords {
		if i[0] == message[0] {
			chosenWords = append(chosenWords, i)
		}
	}

	if len(chosenWords) == 0 {
		chosenWords = []string{spookyWords[rand.Intn(len(spookyWords))]}
	}

	s.ChannelMessageSend(m.ChannelID, "The "+chosenWords[rand.Intn(len(chosenWords))]+" "+message)
}

func LoadSpookyWords() []string {
	read, err := ioutil.ReadFile("spooky.json")
	if err != nil {
		fmt.Println("Oops... Spooky failed.")
		return nil
	}
	var v SpookyWords
	json.Unmarshal(read, &v)
	return v.Words
}
