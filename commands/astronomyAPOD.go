package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	APODApplicationData = discordgo.ApplicationCommandOption{
		Name:        "today",
		Description: "Astronomy Picture of the Day",
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
	APODCommandData class.CommandData
)

type APODOption struct {
	Private bool
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	APODCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func APOD(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	RemoveToken(s, i, i.Member.User.ID)

	margs := APODOption{}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	data := AstronomyClass.APOD()
	height, width := AstronomyClass.GetImageSize(data.URL)

	res := []*discordgo.MessageEmbed{
		{
			Title:       data.Title,
			Description: data.Explanation,
			Color:       0xf2c56b,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Was taken in " + data.Date + " by " + data.Copyright,
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    data.URL,
				Width:  width,
				Height: height,
			},
		},
	}
	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "today")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "today")
		}
	}
}
