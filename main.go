package main

import (
	"discordWeightTracker/bot"
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
	bot.InitDB()
	err := bot.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
	return
}
