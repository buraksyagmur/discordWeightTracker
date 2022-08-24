package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session
var wrongCom bool

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotId = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	slice := strings.Split(m.Content, " ")
	command := slice[0]
	wrongCom = true
	registered := false
	commandSl := strings.Split(command, "")
	if commandSl[0] == "!" {
		if command != "!register" {
			allUsers := findUsers()
			if len(allUsers) < 1 {
				s.ChannelMessageSend(m.ChannelID, "type !register")
				return
			}
			for i := 0; i < len(allUsers); i++ {
				if allUsers[i].username == m.Author.Username {
					registered = true
				} else if (i == len(allUsers)-1 || i < 1) && registered == false {
					s.ChannelMessageSend(m.ChannelID, "type !register")
				}
			}
		}
	}
	var value string
	if len(slice) > 1 {
		value = slice[1]
	}
	if command == "!register" {
		wrongCom = false
		register(s, m, m.Author.Username)
	}
	if command == "!addNew" {
		wrongCom = false
		if value == "" {
			s.ChannelMessageSend(m.ChannelID, "invalid data type !help")
			return
		} else {
			fmt.Println(m.Author.Username)
			addNew(s, m, m.Author.Username, value)
		}
	}
	if command == "!showAll" {
		wrongCom = false
		if value != "" && m.Author.Username == "brksygmr" {
			showAll(s, m, value)
		} else {
			showAll(s, m, m.Author.Username)
		}
	}
	if command == "!deleteLast" {
		wrongCom = false
		deleteLast(s, m, m.Author.Username)
	}
	if command == "!deleteAll" {
		wrongCom = false
		deleteAll(s, m, m.Author.Username)
	}
	if command == "!showSum" {
		wrongCom = false
		showSum(s, m, m.Author.Username)
	}
	if command == "!help" {
		wrongCom = false
		help(s, m, m.Author.Username)
	}
	if command == "!executeOrder66" {
		wrongCom = false
		order66(s, m, m.Author.Username)
	}
	if command == "!showEvery" && m.Author.Username == "brksygmr" {
		wrongCom = false
		showEvery(s, m)
	}
	if wrongCom {
		help(s, m, m.Author.Username)
	}
}
