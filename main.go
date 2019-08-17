package main

import (
	"taskAssistant/bot"
	"taskAssistant/models"
)

func main() {
	models.InitDB("user=postgres dbname=taskassistant password=admin host=127.0.0.1 sslmode=disable")
	bot.Start()
	<-make(chan struct{})
}
