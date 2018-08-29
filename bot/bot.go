package bot

import (
	"fmt"
	"strings"

	commands "../commands"
	config "../config"
	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

// Start initializes de bot
func Start() {
	discordSession, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	botUser, err := discordSession.User("@me")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	config.SetBotID(botUser.ID)

	discordSession.AddHandler(messageHandler)

	err = discordSession.Open()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("The Bot is running")
}

func messageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if strings.HasPrefix(msg.Content, config.Prefix) {
		if msg.Author.ID == config.BotID {
			return
		}
		var rl = config.Roles

		//check if user has permission to interact with bot
		if checkRole(rl, session, msg) == false {
			return
		}
		dcs := commands.DiscordSession{
			Msg:     msg,
			Session: session,
		}

		message := strings.TrimPrefix(msg.Content, config.Prefix)
		messageValues := strings.Split(message, " ")
		command := messageValues[1]
		args := messageValues[2:]
		fmt.Println(args)
		handler := commands.Handler{
			Command: command,
			Args:    args,
		}

		handler.Handle("help", dcs, commands.Help)
		handler.Handle("ping", dcs, commands.Ping)
		handler.Handle("away", dcs, commands.Away)
	}
}
