package commands

import (
	config "../config"
	"github.com/bwmarrin/discordgo"
)

// Help command shows all the commands available
func Help(dcs DiscordSession, args []string) {
	embed := discordgo.MessageEmbed{
		Title:       "Available commands",
		Description: "Here is a list of available commands for the bot.",
		Color:       config.Color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   config.Prefix + " help",
				Value:  "This command. Displays other commands available",
				Inline: false,
			},
		},
	}
	_, _ = dcs.Session.ChannelMessageSendEmbed(dcs.Msg.ChannelID, &embed)
}
