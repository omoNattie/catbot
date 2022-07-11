package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Mods(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == prefix+"ban" {
		s.ChannelMessageSend(m.ChannelID, "not yet UwU!")
	}
}
