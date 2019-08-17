package bot

import (
	"github.com/bwmarrin/discordgo"
)

func Start() {
	const token = "NjExMjM2NDQ1MzMxNTIxNTM2.XVQ4ag.xDlP8ovdgRpfYNaOTzRxAqxXGos"
	discord, err := discordgo.New("Bot " + token)
	errCheck("Error opening connection to discord", err)
	_, err = discord.User("@me")
	errCheck("Error retrieving account", err)

	discord.AddHandler(commandController)
	discord.AddHandler(statusController)

	err = discord.Open()
	errCheck("Error opening connection to discord", err)
}