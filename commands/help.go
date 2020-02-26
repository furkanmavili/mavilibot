package commands

import (
	"bufio"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Help func for !help command
func Help(s *discordgo.Session, m *discordgo.MessageCreate) {

	help := readCommands()
	s.ChannelMessageSend(m.ChannelID, help)

}

func readCommands() string {
	file, err := os.Open("txt/help.txt")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	completeMsg := strings.Join(txtlines, "\n")
	file.Close()
	return completeMsg
}
