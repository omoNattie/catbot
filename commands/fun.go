package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Response struct {
	Url string `json:"url"`
}

func Fun(s *discordgo.Session, m *discordgo.MessageCreate) {
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

		s.ChannelMessageSend(m.ChannelID, image)
		s.ChannelMessageSend(m.ChannelID, "Brought to you by catboys api!")
	}
}
