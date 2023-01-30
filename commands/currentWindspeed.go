package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	CurrentWindspeedApplicationData = discordgo.ApplicationCommandOption{
		Name:        "windspeed",
		Description: "Get current wind speed information of everywhere",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "address",
				Description: "Provide an adress and it will give a weather information about that location!!",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString, // what here :<
				Name:        "units",
				Description: "In what unit you want the bot to use?",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "celsius",
						Value: "celsius",
					},
					{
						Name:  "fahrenheit",
						Value: "fahrenheit",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "private",
				Description: "Should it display only for you or public?",
				Required:    false,
			},
		},
		Type: discordgo.ApplicationCommandOptionSubCommand,
	}
	CurrentWindspeedCommandData class.CommandData
)

type CurrentWindspeedOption struct {
	Address string
	Units   string
	Private bool
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	CurrentWindspeedCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func CurrentWindspeed(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	allow := RemoveToken(s, i, i.Member.User.ID)
	if !allow {
		return
	}

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := CurrentWindspeedOption{
		Units: "celsius",
	}

	if option, ok := optionMap["address"]; ok {
		margs.Address = option.StringValue()
	}
	if option, ok := optionMap["units"]; ok {
		margs.Units = option.StringValue()
	}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	address := MapApi.GetAddress(margs.Address)
	if len(address.ResourceSets) == 0 || address.ResourceSets[0].EstimatedTotal < 1 {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Address not found!", nil, nil))
		if err != nil {
			utils.HandleClientError(s, i, err, "windspeed")
		}
		return
	}
	point := address.ResourceSets[0].Resources[0].Point.Coordinates
	data := WeatherClass.GetCurrentWeather(point[0], point[1], margs.Units)
	addressData := address.ResourceSets[0].Resources[0].Address
	last_update := time.Unix(int64(data.Dt), 0).Format("2006-01-02 15:04:05")
	var unit string
	if margs.Units == "celsius" {
		unit = "meter/sec"
	} else {
		unit = "miles/sec"
	}

	res := []*discordgo.MessageEmbed{
		{
			Title:       addressData.PostalCode + " " + addressData.Locality + ", " + addressData.CountryRegion,
			Description: fmt.Sprintf("**Result:**\n - Current wind speed is **%v %v** and current wind gust is **%v %v**. The direction of the wind currently **%v°**\n\nTo learn more about the data's values, that has been displayed:\n - [%v](https://en.wikipedia.org/wiki/Wind_speed)\n - [deg°](https://en.wikipedia.org/wiki/Meteorology)\n - [Wind gust](https://en.wikipedia.org/wiki/Wind_gust)", data.Wind.Speed, unit, data.Wind.Gust, unit, data.Wind.Deg, unit),
			Color:       0xf2c56b,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    WeatherURLConverter(data.Weather[0].Icon),
				Width:  200,
				Height: 200,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Last updated " + last_update,
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "windspeed")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "windspeed")
		}
	}
}
