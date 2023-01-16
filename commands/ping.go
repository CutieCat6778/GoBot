package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"github.com/bwmarrin/discordgo"
)

var (
	PingApplicationData = discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "It will pong you back :)",
	}
	PingCommandData class.CommandData
)

func init() {

	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	PingCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   0,
		BotPerms:    defaultPerms,
	}
}

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("You don't have any token to use this command. Please wait 6h to retry!\n\n> To view your current tokens, you can easily check with command `/aboutme` and learn more about it!\n> To renew your token faster, just vote us on https://top.gg/bot/1055553353754628197/vote", nil, nil))
	if err != nil {
		utils.HandleServerError(err)
	}
}
