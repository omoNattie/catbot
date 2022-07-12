package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/omoNattie/catbot/commands"
)

func main() {
	//get the token from .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	token := os.Getenv("token")

	// initialize bot
	catbot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// handlers
	catbot.AddHandler(commands.CatBoy)
	catbot.AddHandler(commands.Avatar)
	catbot.AddHandler(commands.Mods)

	catbot.Identify.Intents = discordgo.IntentGuildMessages

	//start the bot itself
	err = catbot.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("UwU! I am awake now, hehe!\nYou know what to do to kill me! (aka press Ctrl+C)")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	catbot.Close()
}
