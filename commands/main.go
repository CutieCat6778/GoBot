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
	commands        = []*discordgo.ApplicationCommand{&PingApplicationData, &MapApplicationData, &WeatherApplicationData}
	commandHandlers = map[string]Command{
		"ping": {
			Execute: Ping,
			Data:    PingCommandData,
		},
		"map": {
			Execute: Map,
			Data:    MapCommandData,
		},
		"current": {
			Execute: CurrentWeather,
			Data:    CurrentWeatherCommandData,
		},
		"weather": {
			Execute: WeatherFunc,
			Data:    WeatherCommandData,
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
