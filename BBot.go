package main

import (
	"BBot/module"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jonas747/discordgo"
)

var randomSource = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	dcSession, err := discordgo.New("Bot " + module.BotConfiguration.Token)
	if err != nil {
		fmt.Println("Error: error appear when creating Discord session")
		return
	}

	peopleResponseRateLimitation := module.NewResponseRateLimitation()

	dcSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if v, ok := module.ChannelsMapping[m.ChannelID]; ok {
			module.GetResponseHandling(v)
			response := module.ResponseHandlingMap[v]
			if messages, ok := response.ByPeople[m.Author.ID]; ok && peopleResponseRateLimitation.Check(m.Author.ID) {
				var i int
				if len(messages) >= 1 {
					i = randomSource.Intn(len(messages))
				} else {
					i = 0
				}
				s.ChannelMessageSend(m.ChannelID, messages[i])
			}
			if messages, ok := response.ByKeyword[m.Content]; ok {
				var i int
				if len(messages) >= 1 {
					i = randomSource.Intn(len(messages))
				} else {
					i = 0
				}
				s.ChannelMessageSend(m.ChannelID, messages[i])
			} else {
				for k, messages := range response.ByKeyword {
					if strings.Contains(strings.ToLower(m.Content), strings.ToLower(k)) {
						var i int
						if len(messages) >= 1 {
							i = randomSource.Intn(len(messages))
						} else {
							i = 0
						}
						s.ChannelMessageSend(m.ChannelID, messages[i])
						break
					}
				}
			}
		}
	})

	err = dcSession.Open()
	if err != nil {
		fmt.Println("Error: error appear when opening connection")
		return
	}
	fmt.Println("Bot is now running. Press CTRL-C to reboot.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc
	dcSession.Close()
	fmt.Println("Bot reboot")
	cmd := exec.Command(os.Args[0])
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Start()
}
