package discordBot

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type user struct {
	username string
	history  string
}
var errMessage string= " type !help for help about commands"
func register(s *discordgo.Session, m *discordgo.MessageCreate, uname string) {
	uName := uname
	history := ""
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", uName)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		s.ChannelMessageSend(m.ChannelID, "you already registered" + errMessage)

	} else {
		stmt, err := db.Prepare("INSERT INTO users (username, history) VALUES (?,?);")
		if err != nil {
			log.Fatal(err)
			s.ChannelMessageSend(m.ChannelID, errMessage)
		}
		defer stmt.Close()
		stmt.Exec(uName, history)
		s.ChannelMessageSend(m.ChannelID, "registered")
	}

}

func showAll(s *discordgo.Session, m *discordgo.MessageCreate, uname string) {
	var usr user
	rows, err := db.Query("SELECT history FROM users WHERE username = ?;", uname)
	if err != nil {
		log.Fatal(err)
		s.ChannelMessageSend(m.ChannelID, errMessage)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&(usr.history))
	}
	slice := strings.Split(usr.history, "%")
	newHistory := strings.Join(slice, "\n")
	if len(slice) > 1 {
		s.ChannelMessageSend(m.ChannelID, newHistory)
	} else {
		if slice[0] != "" {
			s.ChannelMessageSend(m.ChannelID, slice[0]+"you have only one data")
		} else {
			s.ChannelMessageSend(m.ChannelID, "you dont have any data")
		}
	}

}

func addNew(s *discordgo.Session, m *discordgo.MessageCreate, uname string, value string) {
	dt := time.Now()
	slice := strings.Split(value, "")
	first, err := strconv.Atoi(slice[0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "invalid data")
		s.ChannelMessageSend(m.ChannelID, errMessage)
		log.Fatal(err)
		return
	}
	first, err = strconv.Atoi(slice[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "invalid data")
		s.ChannelMessageSend(m.ChannelID, errMessage)
		log.Fatal(err, first)
		return
	}
	var usr user
	rows, err := db.Query("SELECT history FROM users WHERE username = ?;", uname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&(usr.history))
	}
	if usr.history == ""{
	usr.history = dt.Format("2006-02-01") + " " + value
	}else {
		usr.history +="+" +  dt.Format("2006-02-01") + " " + value
	}
	stmt, err := db.Prepare("UPDATE users SET history = ?	WHERE username = ?;")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errMessage)
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(usr.history, uname)
	s.ChannelMessageSend(m.ChannelID, "added")
}

func deleteAll(s *discordgo.Session, m *discordgo.MessageCreate, uname string) {
	stmt, err := db.Prepare("DELETE FROM users WHERE username = ?;")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errMessage)
		log.Fatal(err)
	}
	s.ChannelMessageSend(m.ChannelID, "deleted-all")
	defer stmt.Close()
	stmt.Exec(uname)
}

func deleteLast(s *discordgo.Session, m *discordgo.MessageCreate, uname string) {
	var usr user
	rows, err := db.Query("SELECT history FROM users WHERE username = ?;", uname)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errMessage)
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&(usr.history))
	}
	slice := strings.Split(usr.history, "+")
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
		newHistory := strings.Join(slice, "+")
		stmt, err := db.Prepare("UPDATE users SET history = ?	WHERE username = ?;")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, errMessage)
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec(newHistory, uname)
		s.ChannelMessageSend(m.ChannelID, "deleted-last")
	} else {
		s.ChannelMessageSend(m.ChannelID, "you dont have any data")
		return
	}
}
func showSum(s *discordgo.Session, m *discordgo.MessageCreate, uname string) {
	var usr user
	var lastone string
	var lastsecond string
	rows, err := db.Query("SELECT history FROM users WHERE username = ?;", uname)
	if err != nil {
		log.Fatal(err)
		s.ChannelMessageSend(m.ChannelID, errMessage)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&(usr.history))
	}
	slice := strings.Split(usr.history, "+")
	if len(slice) > 1 {
		lastone = slice[len(slice)-1]
		lastsecond = slice[len(slice)-2]
	} else {
		s.ChannelMessageSend(m.ChannelID, "you dont have enough data for compare")
		return
	}
	sliceforlast := strings.Split(lastone, " ")
	sliceforsecond := strings.Split(lastsecond, " ")
	lastdate, err := time.Parse("2006-02-01", sliceforlast[0])
	if err != nil {
		log.Fatal(err)
		s.ChannelMessageSend(m.ChannelID, errMessage)
	}
	secondlastdate, err := time.Parse("2006-02-01", sliceforsecond[0])
	if err != nil {
		log.Fatal(err)
		s.ChannelMessageSend(m.ChannelID, errMessage)
	}
	subDates := lastdate.Sub(secondlastdate)
	lastWeight, err := strconv.ParseFloat(sliceforlast[1], 64)
	if err == nil {
	}
	secondlastWeight, err := strconv.ParseFloat(sliceforsecond[1], 64)
	if err == nil {
	}
	diffWeight := lastWeight - secondlastWeight
	diffWeightInt := math.Round(diffWeight*100) / 100
	Days := strconv.FormatInt(int64(subDates.Hours()/24), 10)
	WeightStr := fmt.Sprintf("%.2f", diffWeightInt)
	LastMes := "days: " + Days + "\nweight difference: " + WeightStr
	if diffWeightInt > 0 {
		s.ChannelMessageSend(m.ChannelID, LastMes+ " try harder")
	}else if diffWeightInt < 0 {
		s.ChannelMessageSend(m.ChannelID, LastMes+ " well done")
	}else {
		s.ChannelMessageSend(m.ChannelID, LastMes+ " you need a new scale")
	}
}

func help(s *discordgo.Session, m *discordgo.MessageCreate, uname string) {
	helpMes:= "commands:\n!register\n!addNew\n!showSum\n!deleteLast\n!deleteAll\nexample for weight: 31.69"
	s.ChannelMessageSend(m.ChannelID, "Hello "+uname+"!"+"\n"+helpMes)
}