package commands

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Ball func
func Ball(s *discordgo.Session, m *discordgo.MessageCreate) {

	rand.Seed(time.Now().UTC().UnixNano())
	check := rand.Intn(2)
	if check == 0 {
		s.ChannelMessageSend(m.ChannelID, "hayır")
	} else {
		s.ChannelMessageSend(m.ChannelID, "evet")
	}
}
