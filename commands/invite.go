package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Inviter func
func Inviter(s *discordgo.Session, m *discordgo.MessageCreate) error {
	invite := &discordgo.Invite{
		MaxAge:    1,
		MaxUses:   3,
		Temporary: true,
	}
	sonuc, err := s.ChannelInviteCreate(m.ChannelID, *invite)
	if err != nil {
		return err
	}
	s.ChannelMessageSend(m.ChannelID, "discord.gg/"+sonuc.Code)
	s.ChannelMessageSend(m.ChannelID, "1 gün geçerli max 3 kullanımlık geçici davet oluşturuldu.")
	return nil
}
