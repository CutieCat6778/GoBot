package main

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/events"
	"cutiecat6778/discordbot/utils"
	"log"
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
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	s.Identify.Intents = discordgo.IntentGuilds

	s.AddHandler(events.InteractionCreate)
	s.AddHandler(events.Ready)

	err := s.Open()

	if err != nil {
		log.Fatal("Cannot open a session: ", err)
		return
	}

	log.Println("Regsitering commands")

	slashCommands := commands.SlashCommands()

	registeredCommands := make([]*discordgo.ApplicationCommand, len(slashCommands))
	for i, v := range slashCommands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, class.ServerID, v)
		if err != nil {
			utils.SendErrorMessage("Error while resgistering commands! ", err.Error())
			log.Fatal("Error while registering commands: ", err)
		}

		registeredCommands[i] = cmd
	}

	log.Println("Bot is running right now!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	s.Close()

	if *class.RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, class.ServerID, v.ID)
			if err != nil {
				utils.SendErrorMessage("Failed to delete commands ", err.Error())
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
