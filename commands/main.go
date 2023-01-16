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
	commands        = []*discordgo.ApplicationCommand{&HelpApplicationData, &PingApplicationData, &MapApplicationData, &WeatherApplicationData, &AstronomyApplicationData, &ClimateApplicationData, &MeApplicationData}
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
		"aboutme": {
			Execute: Me,
			Data:    MeCommandData,
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
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("You don't have any token to use this command. Please wait 6h to retry!\n\n> To view your current tokens, you can easily check with command `/aboutme` and learn more about it!\nTo renew your token faster, just vote us on https://top.gg/bot/1055553353754628197/vote", nil, nil))
			if err != nil {
				utils.HandleServerError(err)
			}
			return false
		}
	}

	return allow
}
