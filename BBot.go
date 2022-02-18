package main

import (
	"BBot/module"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/jonas747/discordgo"
)

var randomSource = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	dcSession, err := discordgo.New("Bot " + module.BotSetting.Token)
	if err != nil {
		fmt.Println("Error: error appear when creating Discord session")
		return
	}

	dcSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if v, ok := module.ChannelsSetting[m.ChannelID]; ok {
			module.GetResponseSetting(v)
			response := module.ResponseSettingMap[v]
			if messages, ok := response.People[m.Author.ID]; ok {
				var i int
				if len(messages) >= 1 {
					i = randomSource.Intn(len(messages))
				} else {
					i = 0
				}
				s.ChannelMessageSend(m.ChannelID, messages[i])
			}
			if messages, ok := response.KeyWord[m.Content]; ok {
				var i int
				if len(messages) >= 1 {
					i = randomSource.Intn(len(messages))
				} else {
					i = 0
				}
				s.ChannelMessageSend(m.ChannelID, messages[i])
			}
		}
	})

	err = dcSession.Open()
	if err != nil {
		fmt.Println("Error: error appear when opening connection")
		return
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc
	dcSession.Close()
	fmt.Println("Bot reboot")
	cmd := exec.Command(os.Args[0])
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Start()
}
