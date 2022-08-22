package main

import (
	"discordWeightTracker/bot"
	"fmt"
)

func main() {
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
