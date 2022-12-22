package main

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"fmt"
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

func init() {

	slashHandlers := commands.SlashHandlers()

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		g, f := database.FindByServerID(i.GuildID)
		if !f {
			id := database.CreateGuild(i.GuildID)
			log.Println(id)
			g, _ = database.FindByID(id)
		}
		if h, ok := slashHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i, g)
		}
	})
}

func main() {
	s.Identify.Intents = discordgo.IntentGuilds

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		utils.SendLogMessage("The bot is running!")
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

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

	fmt.Println("Bot is running right now!")
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
