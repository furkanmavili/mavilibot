package commands

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var txtlines = quotes()

// Hitler func
func Hitler(s *discordgo.Session, m *discordgo.MessageCreate) error {

	rand.Seed(time.Now().UTC().UnixNano())
	randomquote := txtlines[rand.Intn(len(txtlines)-1)]
	s.ChannelMessageSend(m.ChannelID, randomquote)
	return nil
}

func quotes() []string {
	file, err := os.Open("txt/hitlerquotes.txt")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}
