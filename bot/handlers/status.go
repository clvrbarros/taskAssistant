package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func StatusHandler(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "Your Assistant")
	if err != nil {
		fmt.Printf("Error attempting to set status: %+v", err)
	}
}