package main

import (
	"log"
	"strconv"
	"time"

	"github.com/yanzay/tbot"
)

func main() {
	token := "685029975:AAG5k6THWIqDwcP83in8sY5B2rz-VQpIHcs"
	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/timer {seconds}", timerHandler)
	bot.Handle("/name", "My name is uulnoobBot, please be kind with me")
	bot.Handle("/menu", "this is for menu")
	bot.ListenAndServe()
}

func timerHandler(m *tbot.Message) {
	// m.Vars contains all variables, parsed during routing
	secondsStr := m.Vars["seconds"]
	// Convert string variable to integer seconds value
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		m.Reply("Invalid number of seconds")
		return
	}
	m.Replyf("Timer for %d seconds started", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}
