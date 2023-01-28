package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	AstronomyApplicationData = discordgo.ApplicationCommand{
		Name:        "astronomy",
		Description: "A command group for astronomy commands",
		Options:     []*discordgo.ApplicationCommandOption{&APODApplicationData, &EarthApplicationData},
	}
	AstronomyClass             api.Astronomy
	AstronomyCommandData       class.CommandData
	AstronomySubCommandData    map[string]class.CommandData
	AstronomySubCommandHandler map[string]Command
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
		SubCommandData: AstronomySubCommandData,
	}
	AstronomySubCommandData = map[string]class.CommandData{
		"today": APODCommandData,
		"earth": EarthCommandData,
	}
	AstronomySubCommandHandler = map[string]Command{
		"today": {
			Execute: APOD,
			Data:    APODCommandData,
		},
		"earth": {
			Execute: Earth,
			Data:    EarthCommandData,
		},
	}
}

func AstronomyFunc(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	if h, ok := AstronomySubCommandHandler[options[0].Name]; ok {
		utils.Debug.Println("Subcommand: ", options[0].Name)
		h.Execute(s, i, g)
	}
}
