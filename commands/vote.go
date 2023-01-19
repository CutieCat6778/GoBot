package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	VoteApplicationData = discordgo.ApplicationCommand{
		Name:        "vote",
		Description: "How to vote GeoBot and collect tokens for commands!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "private",
				Description: "Should it display only for you or public?",
				Required:    false,
			},
		},
	}
	VoteCommandData class.CommandData
)

func init() {

	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	VoteCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   0,
		BotPerms:    defaultPerms,
	}
}

func Vote(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {

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

	embeds := []*discordgo.MessageEmbed{
		{
			Title:       "Vote for GeoBot?",
			Description: fmt.Sprintf("**What are the benefits of voting for me?**\n> If you use command `/aboutme`, you will see that you have tokens, because GeoBot is running as a non-profit bot and all of informations that it takes from other APIs are costing us every month. So vote for GeoBot and help your developer team to be motivated and develope more commands :)\n**How can I vote?**\n> To vote for top.gg just [click here](https://top.gg/bot/1055553353754628197/vote)!\n\n**Thank you for support our GeoBot team <3**"),
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Your GeoBot team <3",
				IconURL: s.State.User.AvatarURL(""),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: s.State.User.AvatarURL(""),
			},
			Color: 0xf2c56b,
		},
	}

	if margs.Private {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(embeds, nil))
		if err != nil {
			utils.HandleServerError(err)
		}
		return
	}

	err := s.InteractionRespond(i.Interaction, utils.SendEmbed(embeds, nil))
	if err != nil {
		utils.HandleServerError(err)
	}
}
