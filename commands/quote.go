package commands

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Quote func
func Quote(s *discordgo.Session, m *discordgo.MessageCreate) error {

	rand.Seed(time.Now().UTC().UnixNano())
	quote := allQuotes()
	randomquote := quote[rand.Intn(len(quote)-1)]
	s.ChannelMessageSend(m.ChannelID, randomquote)
	return nil
}

func allQuotes() []string {
	file, err := os.Open("txt/quotes.txt")

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
