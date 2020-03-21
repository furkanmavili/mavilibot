package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Inviter func
func Invite(s *discordgo.Session, m *discordgo.MessageCreate) {
	invite := &discordgo.Invite{
		MaxAge:    1,
		MaxUses:   3,
		Temporary: true,
	}
	sonuc, err := s.ChannelInviteCreate(m.ChannelID, *invite)
	if err != nil {
		fmt.Println("davet oluştururken bir hata oldu")
	}
	s.ChannelMessageSend(m.ChannelID, "discord.gg/"+sonuc.Code)
	s.ChannelMessageSend(m.ChannelID, "1 gün geçerli max 3 kullanımlık geçici davet oluşturuldu.")
}
