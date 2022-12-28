package commands

import (
	"cutiecat6778/discordbot/class"

	"github.com/bwmarrin/discordgo"
)

var (
	ClimateApplicationData = discordgo.ApplicationCommand{
		Name:        "climate",
		Description: "A command group for climate commands",
		Options:     []*discordgo.ApplicationCommandOption{&SeaIceApplicationData, &SeaLevelApplicationData},
	}
	ClimateCommandData       class.CommandData
	ClimateSubCommandData    map[string]class.CommandData
	ClimateSubCommandHandler map[string]Command
)

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)
	ClimateCommandData = class.CommandData{
		Permissions:    defaultPerms,
		Ratelimit:      5000,
		BotPerms:       defaultPerms,
		SubCommandData: ClimateSubCommandData,
	}
	ClimateSubCommandData = map[string]class.CommandData{
		"seaice":   SeaIceCommandData,
		"sealevel": SeaLevelCommandData,
	}
	ClimateSubCommandHandler = map[string]Command{
		"seaice": {
			Execute: SeaIce,
			Data:    SeaIceCommandData,
		},
		"sealevel": {
			Execute: SeaLevel,
			Data:    SeaLevelCommandData,
		},
	}
}

func ClimateFunc(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	if h, ok := ClimateSubCommandHandler[options[0].Name]; ok {
		h.Execute(s, i, g)
	}
}
