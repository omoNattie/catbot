package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Fun(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == prefix+"hi" {
		s.ChannelMessageSend(m.ChannelID, "Hey to you too UwU!")
	}
}
