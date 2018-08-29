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
	var dateString = reDate.FindString(m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	sString := strings.Split(m.Content, " ")
	if sString[0] == commandPrefix {

		//this is to check if member has a role thats allowed to use the bot
		if checkRole(s, m) {
			fmt.Println("yeah boi! got permission")
			if dateString != "" {

				s.ChannelMessageSend(m.ChannelID, "slacker! but its ok")
				if reDate.MatchString(sString[1]) {
					layout := "2006-01-02T15:04:05.000Z"
					fmt.Println("theres a date!?")
					// TODO: add date range check
					t, err := time.Parse(layout, sString[1]+"T19:00:00.000Z")

					if err != nil {
						fmt.Println(err)
					}
					s.ChannelMessageSend(m.ChannelID, "Your attendance for "+sString[1]+" has been confirmed")
					fmt.Println(t)
					fmt.Println(m.Author)
				}
			}
			if dateString == "" {
				s.ChannelMessageSend(m.ChannelID, "oi your missing date. Enter is correct format. Example !afk yyy-mm-dd")
			}
		}

	}

}
