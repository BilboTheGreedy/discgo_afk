package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//regex for date
	reDate := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	if m.Author.ID == s.State.User.ID {
		return
	}

	sString := strings.Split(m.Content, " ")
	if sString[0] == "!afk" {
		s.ChannelMessageSend(m.ChannelID, "slacker! but its ok")
		if reDate.MatchString(sString[1]) {
			fmt.Println("theres a date!?")
		}
	}

}
