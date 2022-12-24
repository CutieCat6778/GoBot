package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	WeatherApplicationData = discordgo.ApplicationCommand{
		Name:        "weather",
		Description: "A command group for weather commands",
		Options: []*discordgo.ApplicationCommandOption{
			&CurrentWeatherApplicationData,
		},
	}
	WeatherClass       api.Weather
	WeatherCommandData class.CommandData
)

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)
	WeatherClass = api.NewWeather()
	WeatherCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func WeatherFunc(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	log.Println(options[0].Name)

	switch options[0].Name {
	case "current":
		CurrentWeather(s, i, g)
	}
}
