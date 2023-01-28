package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	WeatherApplicationData = discordgo.ApplicationCommand{
		Name:        "weather",
		Description: "A command group for weather commands",
		Options: []*discordgo.ApplicationCommandOption{
			&CurrentWeatherApplicationData,
			&CurrentTemperaturApplicationData,
			&CurrentWindspeedApplicationData,
		},
	}
	WeatherClass             api.Weather
	WeatherCommandData       class.CommandData
	WeatherSubCommandData    map[string]class.CommandData
	WeatherSubCommandHandler map[string]Command
)

func WeatherURLConverter(id string) string {
	url := "http://openweathermap.org/img/wn/%v@4x.png"

	url = fmt.Sprintf(url, id)
	return url
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)
	WeatherClass = api.NewWeather()
	WeatherCommandData = class.CommandData{
		Permissions:    defaultPerms,
		Ratelimit:      5000,
		BotPerms:       defaultPerms,
		SubCommandData: WeatherSubCommandData,
	}
	WeatherSubCommandData = map[string]class.CommandData{
		"current":    CurrentWeatherCommandData,
		"temperatur": CurrentTemperaturCommandData,
		"windspeed":  CurrentTemperaturCommandData,
	}
	WeatherSubCommandHandler = map[string]Command{
		"current": {
			Execute: CurrentWeather,
			Data:    CurrentWeatherCommandData,
		},
		"temperatur": {
			Execute: CurrentTemperatur,
			Data:    CurrentTemperaturCommandData,
		},
		"windspeed": {
			Execute: CurrentWindspeed,
			Data:    CurrentWindspeedCommandData,
		},
	}
}

func WeatherFunc(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	if h, ok := WeatherSubCommandHandler[options[0].Name]; ok {
		utils.Debug.Println("Subcommand: ", options[0].Name)
		h.Execute(s, i, g)
	}
}
