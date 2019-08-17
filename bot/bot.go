package bot

import (
	"github.com/bwmarrin/discordgo"
)

func Start() {
	const token = "TOKEN"
	discord, err := discordgo.New("Bot " + token)
	errCheck("Error opening connection to discord", err)
	_, err = discord.User("@me")
	errCheck("Error retrieving account", err)

	discord.AddHandler(commandController)
	discord.AddHandler(statusController)

	err = discord.Open()
	errCheck("Error opening connection to discord", err)
}
