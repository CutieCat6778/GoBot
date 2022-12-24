package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	CurrentWeatherApplicationData = discordgo.ApplicationCommandOption{
		Name:        "current",
		Description: "Get current weather information of everywhere",
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
	CurrentWeatherCommandData class.CommandData
)

type CurrentWeatherOption struct {
	Address string
	Units   string
	Private bool
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	CurrentWeatherCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func CurrentWeather(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	m, allow := database.RemoveToken(i.Member.User.ID)
	if !allow {
		if len(m.MemberID) < 5 {
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Failed to remove user token!", nil, nil))
			if err != nil {
				utils.SendErrorMessage("Problem while trying to send private interaction message: ", err.Error())
				log.Fatal(err)
			}
		} else {
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Your user token is expired, please wait 24h to refresh!", nil, nil))
			if err != nil {
				utils.SendErrorMessage("Problem while trying to send private interaction message: ", err.Error())
				log.Fatal(err)
			}
		}
	}

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := CurrentWeatherOption{
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
	point := address.ResourceSets[0].Resources[0].Point.Coordinates
	data := WeatherClass.GetCurrentWeather(point[0], point[1], margs.Units)
	addressData := address.ResourceSets[0].Resources[0].Address
	name := address.ResourceSets[0].Resources[0].Name
	conficence := address.ResourceSets[0].Resources[0].Confidence
	last_update := time.Unix(int64(data.Dt), 0).Format("2006-01-02 15:04:05")

	res := []*discordgo.MessageEmbed{
		{
			Title:       addressData.PostalCode + " " + addressData.Locality + ", " + addressData.CountryRegion,
			Description: fmt.Sprintf("Current weather condition is **%v**, it can also be described as **%v**\n\n**Detailed address information**\n%v\n**Confidence**\n %v", data.Weather[0].Main, data.Weather[0].Description, name, conficence),
			Color:       0xf2c56b,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    api.URLConverter(data.Weather[0].Icon),
				Width:  200,
				Height: 200,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Last updated " + last_update,
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Loading...", res, nil))

		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendInteractionMessage("Loading...", res, nil))

		if err != nil {
			log.Fatal(err)
		}
	}
}
