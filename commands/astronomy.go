package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"

	"github.com/bwmarrin/discordgo"
)

var (
	AstronomyApplicationData = discordgo.ApplicationCommand{
		Name:        "astronomy",
		Description: "A command group for astronomy commands",
		Options:     []*discordgo.ApplicationCommandOption{&APODApplicationData},
	}
	AstronomyClass       api.Astronomy
	AstronomyCommandData class.CommandData
	SubCommandData       map[string]class.CommandData
	SubCommandHandler    map[string]Command
)

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)
	AstronomyClass = api.NewAstronomy()
	AstronomyCommandData = class.CommandData{
		Permissions:    defaultPerms,
		Ratelimit:      5000,
		BotPerms:       defaultPerms,
		SubCommandData: SubCommandData,
	}
	SubCommandData = map[string]class.CommandData{
		"today": APODCommandData,
	}
	SubCommandHandler = map[string]Command{
		"today": {
			Execute: APOD,
			Data:    APODCommandData,
		},
	}
}

func AstronomyFunc(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	if h, ok := SubCommandHandler[options[0].Name]; ok {
		h.Execute(s, i, g)
	}
}
