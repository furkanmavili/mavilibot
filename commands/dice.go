package commands

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Dice fun
func Dice(s *discordgo.Session, m *discordgo.MessageCreate) error {

	rand.Seed(time.Now().UTC().UnixNano())
	check := 1 + rand.Intn(6)
	path := "images/dice" + strconv.Itoa(check) + ".png"
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	fileReader := io.Reader(f)

	file := &discordgo.File{Name: path, Reader: fileReader}

	msg := &discordgo.MessageSend{File: file}

	fmt.Printf("%+v", msg)
	_, err = s.ChannelMessageSendComplex(m.ChannelID, msg)
	if err != nil {
		return err
		s.ChannelMessageSend(m.ChannelID, "something went horribly wrong!")
	}
	return nil
}
