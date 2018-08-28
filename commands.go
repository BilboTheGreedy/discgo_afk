package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

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
			layout := "2006-01-02T15:04:05.000Z"
			fmt.Println("theres a date!?")
			t, err := time.Parse(layout, sString[1]+"T19:00:00.000Z")

			if err != nil {
				fmt.Println(err)
			}
			s.ChannelMessageSend(m.ChannelID, "Your attendance for "+sString[1]+" has been confirmed")
			fmt.Println(t)
		}
	}

}
