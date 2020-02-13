package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := `	Commands
	!delete - delete last message(not finished yet)
	!8ball <question> - answering yes-no question
	!invite - creates guild invite(not finished yet)
	!fuhrer - shares hitler quotes (this command just for my nazi brother)
	!delikan - my favourite song link
	!dice - `
	s.ChannelMessageSend(m.ChannelID, message)

}
