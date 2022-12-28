package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"log"

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

	margs := SeaIceOption{}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	res := []*discordgo.MessageEmbed{
		{
			Title:       "Annual Arctic Sea Ice Minimum Area 1979-2022",
			Color:       0xf2c56b,
			Description: "Satellite-based passive microwave images of the sea ice have provided a reliable tool for continuously monitoring changes in the Arctic ice since 1979. Every summer the Arctic ice cap melts down to what scientists call its \"minimum\" before colder weather begins to cause ice cover to increase. This graph displays the area of the minimum sea ice coverage each year from 1979 through 2022. In 2022, the Arctic minimum sea ice covered an area of 4.16 million square kilometers (1.6 million square miles).\n\nThis visualization shows the expanse of the annual minimum Arctic sea ice for each year from 1979 through 2022 as derived from passive microwave data.\n\n[Video link](https://svs.gsfc.nasa.gov/vis/a000000/a005000/a005036/sea_ice_min_w_graph_2022_1080p30.mp4)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Nasa Scientific Visualization Studio",
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(res, nil))

		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			log.Fatal(err)
		}
	}
}
