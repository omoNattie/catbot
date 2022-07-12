package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Response struct {
	Url string `json:"url"`
}

func CatBoy(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == prefix+"catboy" {
		response, err := http.Get("https://api.catboys.com/img/random")
		if err != nil {
			fmt.Print(err)
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var catboy Response
		json.Unmarshal(responseData, &catboy)

		fmt.Println(catboy.Url)
		image := catboy.Url

		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x00ff00,
			Image: &discordgo.MessageEmbedImage{
				URL: image,
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Brought to you by the catboys.com api!",
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

func Avatar(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == prefix+"avatar" {
		embed := &discordgo.MessageEmbed{
			Color: 0x00ff00,
			Image: &discordgo.MessageEmbedImage{
				URL: m.Author.AvatarURL("auto"),
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Your avatar!",
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
