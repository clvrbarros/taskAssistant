package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	h "taskAssistant/bot/handlers"
)

const (
	HELLO = 1
	ADDTASK = 2
	LISTTASK = 3
)

func commandController (discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}
	content := strings.ToUpper(message.Content)
	command := decodeCommand(content)
	switch command {
	case HELLO:
		h.HelloHandler(discord, message)
	case ADDTASK:
		h.AddTaskHandler(discord, message)
	case LISTTASK:
		h.ListTaskHandler(discord, message)
	default:
		//Do nothing
	}
}

func statusController (discord *discordgo.Session, ready *discordgo.Ready) {
	h.StatusHandler(discord, ready)
	servers := discord.State.Guilds
	fmt.Printf("Task Assistant has started on %d servers\n", len(servers))
}

func decodeCommand(c string) int {
	if c == "HEY ASSISTANT" || c == "HEY, ASSISTANT" {
		return HELLO
	} else if strings.HasPrefix(c,"!TASK ADD ") {
		return ADDTASK
	} else if strings.HasPrefix(c, "!TASK LIST") {
		return LISTTASK
	} else {
		return -1
	}
}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
	}
}