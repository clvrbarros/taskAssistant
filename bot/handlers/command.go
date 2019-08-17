package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"taskAssistant/models"
)

func HelloHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = discord.ChannelMessageSend(message.ChannelID, "**Hello "+message.Author.Mention()+". For add a task: "+
		"!task add Description. For view your tasks: !task list. Need help? !task help**")
}

func AddTaskHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	task := message.Content[10:]
	fmt.Println(task, message.Author.ID)
	err := models.AddTask(task, message.Author.ID)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = discord.ChannelMessageSend(message.ChannelID, message.Author.Mention() + " your task has been added!")
}

func ListTaskHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	tasks, err := models.TasksById(message.Author.ID)
	if err != nil {
		log.Println(err)
	}
	result := ""
	for _, task := range tasks {
		a := fmt.Sprintf("%s - Created in %v\n",task.Name,task.Created.Format("01/02/2006"))
		result = result + a
	}
	response := discordgo.MessageEmbed{
		Title:       "Tasks",
		Description: "**These are your tasks:**\n" + result,
		Color:       0x65f442,
	}
	_, _ = discord.ChannelMessageSendEmbed(message.ChannelID, &response)
}