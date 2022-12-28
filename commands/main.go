package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds)
	Data    class.CommandData
}

var (
	commands        = []*discordgo.ApplicationCommand{&PingApplicationData, &MapApplicationData, &WeatherApplicationData, &AstronomyApplicationData}
	commandHandlers = map[string]Command{
		"ping": {
			Execute: Ping,
			Data:    PingCommandData,
		},
		"map": {
			Execute: Map,
			Data:    MapCommandData,
		},
		"weather": {
			Execute: WeatherFunc,
			Data:    WeatherCommandData,
		},
		"astronomy": {
			Execute: AstronomyFunc,
			Data:    AstronomyCommandData,
		},
	}
	MapApi api.Map = api.NewMap()
)

func SlashCommands() []*discordgo.ApplicationCommand {
	return commands
}

func SlashHandlers() map[string]Command {
	return commandHandlers
}
