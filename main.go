package main

import (
	"discordBot/bot"
	"fmt"
)

func main() {
	// discord, err := discordgo.New("Bot" + "OTkxMTY2MzYzMTE4ODY2NjEz.GIy9zU.idtcY2mb7W0hYIbMOxtw1muXdo9V6SyRz864jg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// user, err := discord.User("@me")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	discordBot.InitDB()
	err := discordBot.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	discordBot.Start()

	<-make(chan struct{})
	return
}
