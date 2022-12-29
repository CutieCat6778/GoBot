package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	CO2ApplicationData = discordgo.ApplicationCommandOption{
		Name:        "co2airs",
		Description: "20 years of AIRS Global Carbon Dioxide (CO₂) measurements (2002- March 2022)",
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
	CO2CommandData class.CommandData
)

type CO2Option struct {
	Private bool
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	CO2CommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func CO2(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := CO2Option{}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	res := []*discordgo.MessageEmbed{
		{
			Title:       "20 years of AIRS Global Carbon Dioxide (CO₂) measurements (2002- March 2022)",
			Color:       0xf2c56b,
			Description: "This data visualization shows the global distribution and variation of the concentration of mid-tropospheric carbon dioxide observed by the Atmospheric Infrared Sounder (AIRS) on the NASA Aqua spacecraft over a 20 year timespan. One obvious feature that we see in the data is a continual increase in carbon dioxide with time, as seen in the shift in the color of the map from light yellow towards red as time progresses. Another feature is the seasonal variation of carbon dioxide in the northern hemisphere, which is governed by the growth cycle of plants. This can be seen as a pulsing in the colors, with a shift towards lighter colors starting in April/May each year and a shift towards red as the end of each growing season passes into winter. The seasonal cycle is more pronounced in the northern hemisphere than the southern hemisphere, since the majority of the land mass is in the north.\n\nThe visualization includes a data-driven spatial map of global carbon dioxide and a timeline on the bottom. The timeline showcases the monthly timestep and is paired with the adjusted carbon dioxide value. Areas where the air pressure is less than 750mB (areas of high-altitude) have been marked in the visualization as low data quality (striped) areas. This entry offers two versions of low data quality (stiped) areas. One version includes striped regions as they are calculated on data values and the second version features striped regions below 60 South.\n\n[Resources](https://svs.gsfc.nasa.gov/4990)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Nasa | Atmospheric Infrared Sounder (AIRS)",
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://cdn.thinh.tech/climate/co2.gif",
				Width:  800,
				Height: 450,
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "co2airs")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "co2airs")
		}
	}
}
