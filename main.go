package main

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/events"
	"cutiecat6778/discordbot/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + class.Token)
	if err != nil {
		utils.HandleServerError(err)
	}
}

func main() {
	s.Identify.Intents = discordgo.IntentGuilds
	s.Identify.Intents = discordgo.IntentGuildMembers

	s.AddHandler(events.InteractionCreate)
	s.AddHandler(events.Ready)

	err := s.Open()

	if err != nil {
		utils.HandleServerError(err)
	}

	utils.HandleDebugMessage("Registering commands")

	slashCommands := commands.SlashCommands()

	registeredCommands := make([]*discordgo.ApplicationCommand, len(slashCommands))
	for i, v := range slashCommands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			utils.HandleServerError(err)
		}

		registeredCommands[i] = cmd
	}

	utils.HandleDebugMessage("Bot is running right now!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	s.Close()

	if *class.RemoveCommands {
		utils.HandleDebugMessage("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
			if err != nil {
				utils.HandleServerError(err)
			}
		}
	}

	utils.HandleDebugMessage("Gracefully shutting down.")
}
