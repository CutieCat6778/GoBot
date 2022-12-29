package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	SeaIceApplicationData = discordgo.ApplicationCommandOption{
		Name:        "seaice",
		Description: "Visualization of sea ice from 1979 to 2022",
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
	SeaIceCommandData class.CommandData
)

type SeaIceOption struct {
	Private bool
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	SeaIceCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func SeaIce(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := SeaIceOption{}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	res := []*discordgo.MessageEmbed{
		{
			Title:       "Annual Arctic Sea Ice Minimum Area 1979-2022",
			Color:       0xf2c56b,
			Description: "Satellite-based passive microwave images of the sea ice have provided a reliable tool for continuously monitoring changes in the Arctic ice since 1979. Every summer the Arctic ice cap melts down to what scientists call its \"minimum\" before colder weather begins to cause ice cover to increase. This graph displays the area of the minimum sea ice coverage each year from 1979 through 2022. In 2022, the Arctic minimum sea ice covered an area of 4.16 million square kilometers (1.6 million square miles).\n\nThis visualization shows the expanse of the annual minimum Arctic sea ice for each year from 1979 through 2022 as derived from passive microwave data.\n\n[Resources](https://svs.gsfc.nasa.gov/5036)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Nasa Scientific Visualization Studio",
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://cdn.thinh.tech/climate/seaice.gif",
				Width:  800,
				Height: 450,
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "seaice")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "seaice")
		}
	}
}
