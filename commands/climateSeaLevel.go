package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	SeaLevelApplicationData = discordgo.ApplicationCommandOption{
		Name:        "sealevel",
		Description: "Visualization of Sea Level in year 2022",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "location",
				Description: "Location of the visualization",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Value: "s-usa",
						Name:  "Southeast United States",
					},
					{
						Value: "n-eu",
						Name:  "Northern Europe",
					},
					{
						Value: "bz",
						Name:  "Amazon Delta",
					},
					{
						Value: "s-as",
						Name:  "Southeast Asia",
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
	SeaLevelCommandData class.CommandData
)

type SeaLevelOption struct {
	Location string
	Private  bool
}

var (
	defaultPerms   int64 = discordgo.PermissionSendMessages
	SeaLevelScroll class.SeaLevelScroll
)

func init() {

	SeaLevelScroll = *class.NewSeaLevelScroll()

	SeaLevelCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func SeaLevel(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := SeaLevelOption{}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}
	if option, ok := optionMap["location"]; ok {
		margs.Location = option.StringValue()
	}

	url := URLResolver(margs.Location, 0)

	log.Println(url)

	height, width := AstronomyClass.GetImageSize(url)
	SeaLevelScroll.Register(i.Member.User.ID, margs.Location)

	embed := []*discordgo.MessageEmbed{
		{
			Title:       "Sea level prediction",
			Color:       0xf2c56b,
			Description: "Recent satellite observations have detected that the Greenland and Antarctic ice sheets are losing ice. Even a partial loss of these ice sheets would cause a 1-meter (3-foot) rise. If lost completely, both ice sheets contain enough water to raise sea level by 66 meters (217 feet).\n\nThis visualization shows the effect on coastal regions for each meter of sea level rise, up to 6 meters (19.7 feet). Land that would be covered in water is shaded red.\n\n[Resources](https://climate.nasa.gov/interactives/climate-time-machine)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Center for Remote Sensing of Ice Sheets",
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    url,
				Width:  width,
				Height: height,
			},
		},
	}

	if margs.Private {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: embed,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Label:    "←",
								Style:    discordgo.PrimaryButton,
								Disabled: true,
								CustomID: "sealevel_left0",
							},
							discordgo.Button{
								Label:    "→",
								Style:    discordgo.PrimaryButton,
								Disabled: false,
								CustomID: "sealevel_right0",
							},
						},
					},
				},
			},
		})
		if err != nil {
			utils.HandleClientError(s, i, err, "sealevel")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: embed,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Label:    "←",
								Style:    discordgo.PrimaryButton,
								Disabled: true,
								CustomID: "sealevel_left0",
							},
							discordgo.Button{
								Label:    "→",
								Style:    discordgo.PrimaryButton,
								Disabled: false,
								CustomID: "sealevel_right0",
							},
						},
					},
				},
			},
		})

		if err != nil {
			utils.HandleClientError(s, i, err, "sealevel")
		}
	}
}

func URLResolver(location string, num int64) string {
	switch location {
	case "s-usa":
		return fmt.Sprintf("https://cdn.thinh.tech/seaLevel/seaLevel6_US_%v.jpeg", num)
	case "n-eu":
		return fmt.Sprintf("https://cdn.thinh.tech/seaLevel/seaLevel6_europe_%v.jpeg", num)
	case "bz":
		return fmt.Sprintf("https://cdn.thinh.tech/seaLevel/seaLevel6_amazon_%v.jpeg", num)
	case "s-as":
		return fmt.Sprintf("https://cdn.thinh.tech/seaLevel/seaLevel6_asia_%v.jpeg", num)
	default:
		return fmt.Sprintf("https://cdn.thinh.tech/seaLevel/seaLevel6_US_%v.jpeg", num)
	}
}
