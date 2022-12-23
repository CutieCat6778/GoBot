package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	MapApplicationData = discordgo.ApplicationCommand{
		Name:        "map",
		Description: "Get a map of a location!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "address",
				Description: "Provide an adress and it will give an image about that!",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "zoom",
				Description: "How zoomed should the map be?",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString, // what here :<
				Name:        "type",
				Description: "What type of the map you want?",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "roadmap",
						Value: "roadmap",
					},
					{
						Name:  "satellite",
						Value: "satellite",
					},
					{
						Name:  "hybrid",
						Value: "hybrid",
					},
					{
						Name:  "terrain",
						Value: "terrain",
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
	}
	MapCommandData class.CommandData
)

type Option struct {
	Address string
	Zoom    int64
	Private bool
	Type    string
}

func init() {

	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	MapCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   0,
		BotPerms:    defaultPerms,
	}
}

func Map(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := Option{
		Zoom: 15,
	}

	if option, ok := optionMap["address"]; ok {
		margs.Address = option.StringValue()
	}
	if option, ok := optionMap["zoom"]; ok {
		margs.Zoom = option.IntValue()
	}
	if option, ok := optionMap["type"]; ok {
		margs.Type = option.StringValue()
	}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	if margs.Zoom >= 20 {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("The zoom values can only be from 1-19!", nil, nil))

		if err != nil {
			log.Fatal(err)
		}
	}

	reader := Api.GetAdress(margs.Address, margs.Zoom, margs.Type)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Files: []*discordgo.File{
				{Name: "image.png", Reader: reader},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
