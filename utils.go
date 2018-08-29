package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func checkRole(s *discordgo.Session, m *discordgo.MessageCreate) bool {

	d, err := s.Channel(m.ChannelID)
	if err != nil {
		return false
	}

	g, err := s.Guild(d.GuildID)
	if err != nil {
		return false
	}

	member, err := s.GuildMember(g.ID, m.Author.ID)
	if err != nil {
		return false
	}

	memberRoles := member.Roles

	for _, role := range g.Roles {
		for _, roleID := range memberRoles {
			if role.ID == roleID {
				for _, allowedRole := range allowedRoles {
					if strings.ToLower(role.Name) == strings.ToLower(allowedRole) {
						fmt.Println(role.Name)
						return true

					}
				}
			}
		}
	}
	return false
}
