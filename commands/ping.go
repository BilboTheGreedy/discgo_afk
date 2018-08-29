package commands

// Ping command responds with a "pong" text
func Ping(dcs DiscordSession, args []string) {
	_, _ = dcs.Session.ChannelMessageSend(dcs.Msg.ChannelID, "pong!")
}
