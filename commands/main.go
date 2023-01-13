package commands

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds)
	Data    class.CommandData
}

var (
	commands        = []*discordgo.ApplicationCommand{&HelpApplicationData, &PingApplicationData, &MapApplicationData, &WeatherApplicationData, &AstronomyApplicationData, &ClimateApplicationData}
	commandHandlers = map[string]Command{
		"help": {
			Execute: Help,
			Data:    HelpCommandData,
		},
		"ping": {
			Execute: Ping,
			Data:    PingCommandData,
		},
		"map": {
			Execute: Map,
			Data:    MapCommandData,
		},
		"weather": {
			Execute: WeatherFunc,
			Data:    WeatherCommandData,
		},
		"astronomy": {
			Execute: AstronomyFunc,
			Data:    AstronomyCommandData,
		},
		"climate": {
			Execute: ClimateFunc,
			Data:    ClimateCommandData,
		},
	}
	MapApi api.Map = api.NewMap()
)

func SlashCommands() []*discordgo.ApplicationCommand {
	return commands
}

func SlashHandlers() map[string]Command {
	return commandHandlers
}

func RemoveToken(s *discordgo.Session, i *discordgo.InteractionCreate, id string) bool {
	m, allow := database.RemoveToken(i.Member.User.ID)

	utils.HandleDebugMessage(m.MemberID, m.Tokens, allow)

	if !allow {
		if len(m.MemberID) < 5 {
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Error while trying to run this command, please contact support!", nil, nil))
			if err != nil {
				utils.HandleServerError(err)
			}
			return false
		} else {
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("You don't have any token to use this command. Please wait 6h to retry!", nil, nil))
			if err != nil {
				utils.HandleServerError(err)
			}
			return false
		}
	}

	return allow
}
