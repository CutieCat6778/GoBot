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
	s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Pong!", nil, nil))
}
