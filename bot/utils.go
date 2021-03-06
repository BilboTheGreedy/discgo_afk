package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func checkRole(rl []string, s *discordgo.Session, m *discordgo.MessageCreate) bool {

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
				for _, r := range rl {
					if strings.ToLower(role.Name) == strings.ToLower(r) {
						fmt.Println(role.Name)
						return true

					}
				}
			}
		}
	}
	return false
}
