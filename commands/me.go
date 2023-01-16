package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

var (
	MeApplicationData = discordgo.ApplicationCommand{
		Name:        "aboutme",
		Description: "Get personal information about you!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "private",
				Description: "Should it display only for you or public?",
				Required:    false,
			},
		},
	}
	MeCommandData class.CommandData
)

func init() {

	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	MeCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   0,
		BotPerms:    defaultPerms,
	}
}

type MeStruct struct {
	Private bool
}

func Me(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	m, f := database.FindUserByMemberID(i.Member.User.ID)
	if !f {
		utils.HandleClientError(s, i, errors.New("user not found, while querying in database"), "aboutme")
		return
	}

	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := MeStruct{
		Private: true,
	}

	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	embed := []*discordgo.MessageEmbed{
		{
			Title:       CapitalizeTitle(i.Member.User.Username),
			Description: fmt.Sprintf("**Joined GeoBot at** \n - %v\n**Current token** \n - %v\n**Tokens can be renew in** \n - %v minutes\n", time.Unix(m.CreatedAt, 0).Format("2006-01-02 15:02"), m.Tokens, (time.Now().Unix()-m.LastRefreshed)/1000),
			Footer: &discordgo.MessageEmbedFooter{
				Text:    fmt.Sprintf("Request by %v", i.Member.User.Username),
				IconURL: i.Member.User.AvatarURL(""),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: i.Member.User.AvatarURL(""),
			},
			Color: 0xf2c56b,
		},
	}

	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(embed, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "aboutme")
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(embed, nil))

		if err != nil {
			utils.HandleClientError(s, i, err, "aboutme")
		}
	}
}
