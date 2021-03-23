package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		msgHead string
		arg     string
		err     error
	)

	if m.Author.ID == s.State.User.ID || len(m.Content) < 2 || m.Content[2:3] != " " {
		return
	}

	// using janky command code for now

	msgHead = m.Content[:2]

	if msgHead == "!n" {
		if m.Content == msgHead {
			err = errors.New("arg not provided")
		} else {
			arg = strings.TrimSpace(m.Content[3:])

			_, err = s.ChannelMessageSend(m.ChannelID, "Nothing can replace "+arg)
		}

		if err != nil {
			fmt.Println("failed to send message", err)
		}

	}
}
