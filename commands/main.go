package commands

import (
	"cutiecat6778/discordbot/class"

	"github.com/bwmarrin/discordgo"
)

var (
	commands        = []*discordgo.ApplicationCommand{&PingApplicationData}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds){
		"ping": Ping,
	}
)

func SlashCommands() []*discordgo.ApplicationCommand {
	return commands
}

func SlashHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	return commandHandlers
}
