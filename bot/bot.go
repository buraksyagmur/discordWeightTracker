package discordBot

import (
	"fmt" //to print errors
	"log"
	"strings"

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

func Start() {

	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")
	//Handlinf error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Storing our id from u to BotId .
	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
	//If we message ping to our bot in our discord it will return us pong .
	slice := strings.Split(m.Content, " ")
	command := slice[0]
	if command != "!register" {
		rows, err := db.Query("SELECT * FROM users WHERE username = ?", m.Author.Username)
		if err != nil {
			log.Fatal(err)
		}
		if !rows.Next() {
			s.ChannelMessageSend(m.ChannelID, "type !register or !help")
			return
		}
	}
	var value string
	if len(slice) > 1 {
		value = slice[1]
	}
	if command == "!register" {
		register(s, m, m.Author.Username)
	}
	if command == "!addNew" {
		if value == ""{
			s.ChannelMessageSend(m.ChannelID, "invalid data type !help")
			return
		}
		addNew(s, m, m.Author.Username, value)
	}
	if command == "!showAll" {
		showAll(s, m, m.Author.Username)
	}
	if command == "!deleteLast" {
		deleteLast(s, m, m.Author.Username)
	}
	if command == "!deleteAll" {
		deleteAll(s, m, m.Author.Username)
	}
	if command == "!showSum" {
		showSum(s, m, m.Author.Username)
	}
	if command == "!help" {
		help(s, m, m.Author.Username)
	}
}
