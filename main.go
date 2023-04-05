package main

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/events"
	"cutiecat6778/discordbot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

var (
	s       *discordgo.Session
	intents discordgo.Intent
)

func init() {
	var err error
	s, err = discordgo.New("Bot " + class.Token)
	if err != nil {
		utils.HandleServerError(err)
	}
	intents = discordgo.IntentGuilds | discordgo.IntentGuildMembers
}

func main() {
	s.Identify.Intents = discordgo.MakeIntent(intents)

	s.AddHandler(events.InteractionCreate)
	s.AddHandler(events.GuildDelete)
	s.AddHandler(events.GuildCreate)
	s.AddHandler(events.Ready)

	err := s.Open()

	if err != nil {
		utils.HandleServerError(err)
	}

	utils.HandleDebugMessage("Registering commands")

	if class.LOCAL {
		slashCommands := commands.SlashCommands()
		registeredCommands := make([]*discordgo.ApplicationCommand, len(slashCommands))
		for i, v := range slashCommands {
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, class.ServerID, v)
			if err != nil {
				utils.HandleServerError(err)
			}

			registeredCommands[i] = cmd
		}
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	utils.HandleDebugMessage("Gracefully shutting down.")
}
