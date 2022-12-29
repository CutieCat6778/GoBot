package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	GlobalTempApplicationData = discordgo.ApplicationCommandOption{
		Name:        "globaltemperatur",
		Description: "Global Temperature Anomalies from 1880 to 2021",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "private",
				Description: "Should it display only for you or public?",
				Required:    false,
			},
		},
		Type: discordgo.ApplicationCommandOptionSubCommand,
	}
	GlobalTempCommandData class.CommandData
)

type GlobalTempOption struct {
	Private bool
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	GlobalTempCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func GlobalTemp(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := GlobalTempOption{}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	res := []*discordgo.MessageEmbed{
		{
			Title:       "Global Temperature Anomalies from 1880 to 2021",
			Color:       0xf2c56b,
			Description: "Earth's global average surface temperature in 2021 tied with 2018 as the sixth warmest on record, according to independent analyses done by NASA and NOAA.\n\nContinuing the planet's long-term warming trend, global temperatures in 2021 were 1.5 degrees Fahrenheit (or 0.85 degrees Celsius) above the average for NASA's baseline period, according to scientists at NASA's Goddard Institute for Space Studies (GISS) in New York.\n\nCollectively, the past eight years are the top eight warmest years since modern record keeping began in 1880. This annual temperature data makes up the global temperature record - and it's how scientists know that the planet is warming.\n\nGISS is a NASA laboratory managed by the Earth Sciences Division of the agency's Goddard Space Flight Center in Greenbelt, Maryland. The laboratory is affiliated with Columbia University's Earth Institute and School of Engineering and Applied Science in New York.\n\nFor more information about NASA's Earth science missions, visit: [https://www.nasa.gov/earth](https://www.nasa.gov/earth)\n\n[Resources](https://svs.gsfc.nasa.gov/4964)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "NASA/GISS | Nasa Scientific Visualization Studio",
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://cdn.thinh.tech/climate/temp.gif",
				Width:  800,
				Height: 450,
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "globaltemperatur")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "globaltemperatur")
		}
	}
}
