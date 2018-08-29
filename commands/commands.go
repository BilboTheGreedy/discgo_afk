package commands

import (
	"github.com/bwmarrin/discordgo"
)

type commandFunc func(DiscordSession, []string)

// Handler type is the message type received
type Handler struct {
	Command string
	Args    []string
}

// DiscordSession type eases the construction of a command function
type DiscordSession struct {
	Msg     *discordgo.MessageCreate
	Session *discordgo.Session
}

// Handle functions processes the request
func (h Handler) Handle(command string, dcs DiscordSession, fn commandFunc) {
	if string(h.Command) == command {
		fn(dcs, h.Args)
	}
	return
}
