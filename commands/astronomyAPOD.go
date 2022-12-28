package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"log"

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
	Address string
	Units   string
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
			log.Fatal(err)
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(res, nil))

		if err != nil {
			log.Fatal(err)
		}
	}
}
