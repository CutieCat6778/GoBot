package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	HelpApplicationData discordgo.ApplicationCommand
	HelpCommandData     class.CommandData
)

func init() {

	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)
	HelpApplicationData = discordgo.ApplicationCommand{
		Name:        "help",
		Description: "The complete guide for you to understand the bot even better!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "command",
				Description: "Select a command that you want to search for!",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "help",
						Value: "help",
					},
					{
						Name:  "astronomy",
						Value: "astronomy",
					},
					{
						Name:  "climate",
						Value: "climate",
					},
					{
						Name:  "weather",
						Value: "weather",
					},
					{
						Name:  "map",
						Value: "map",
					},
					{
						Name:  "ping",
						Value: "ping",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "sub-command",
				Description: "Select a command that you want to search for!",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "astronomy today",
						Value: "today",
					},
					{
						Name:  "climate co2airs",
						Value: "co2airs",
					},
					{
						Name:  "climate sealevel",
						Value: "sealevel",
					},
					{
						Name:  "climate icelevel",
						Value: "icelevel",
					},
					{
						Name:  "climate globaltemperatur",
						Value: "globaltemperatur",
					},
					{
						Name:  "weather current",
						Value: "current",
					},
					{
						Name:  "weather windspeed",
						Value: "windspeed",
					},
					{
						Name:  "weather temperatur",
						Value: "temperatur",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "public",
				Description: "Should it display only for you or public?",
				Required:    false,
			},
		},
	}

	HelpCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   0,
		BotPerms:    defaultPerms,
	}
}

type HelpOption struct {
	Command    string
	SubCommand string
	Public     bool
}

func Help(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := HelpOption{}

	if option, ok := optionMap["command"]; ok {
		value := option.StringValue()
		margs.Command = value
	}
	if option, ok := optionMap["sub-command"]; ok {
		value := option.StringValue()
		margs.SubCommand = value
	}
	if option, ok := optionMap["public"]; ok {
		value := option.BoolValue()
		margs.Public = value
	}

	var embed []*discordgo.MessageEmbed

	if len(margs.Command) > 0 && len(margs.SubCommand) > 0 {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Cannot use 2 option at the same time, please only choose **command** or **sub-command**! You **can't** use both of them!", nil, nil))
		if err != nil {
			utils.HandleClientError(s, i, err, "help")
			return
		}
		return
	}

	if len(margs.Command) > 0 {
		slashCommands := map[string]Command{
			"help": {
				Data: HelpCommandData,
			},
			"ping": {
				Data: PingCommandData,
			},
			"map": {
				Data: MapCommandData,
			},
			"weather": {
				Data: WeatherCommandData,
			},
			"astronomy": {
				Data: AstronomyCommandData,
			},
			"climate": {
				Data: ClimateCommandData,
			},
		}
		commandData := slashCommands[margs.Command]
		slashData, f := findSlashData(margs.Command)
		if !f {
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Unable to find that command name, please report this with **`Problem`** button!", nil, nil))
			if err != nil {
				utils.HandleClientError(s, i, err, "help")
				return
			}
			return
		}

		_, found := utils.CommandBlock.Get(margs.Command)

		embed = []*discordgo.MessageEmbed{
			{
				Color:       0xf2c56b,
				Title:       CapitalizeTitle(slashData.Name),
				Description: slashData.Description + fmt.Sprintf("\n\n**Arguments**\n%v", ArgumentToString(slashData.Options)),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Slowdown",
						Value:  fmt.Sprintf("%v ms", commandData.Data.Ratelimit),
						Inline: true,
					},
					{
						Name:   "Disabled",
						Value:  fmt.Sprintf("%v", found),
						Inline: true,
					},
				},
			},
		}
	} else if len(margs.SubCommand) > 0 {
		slashCommands := map[string]class.CommandData{
			"today":            APODCommandData,
			"seaice":           SeaIceCommandData,
			"sealevel":         SeaLevelCommandData,
			"co2airs":          CO2CommandData,
			"globaltemperatur": GlobalTempCommandData,
			"current":          CurrentWeatherCommandData,
			"temperatur":       CurrentTemperaturCommandData,
			"windspeed":        CurrentTemperaturCommandData,
		}
		commandData := slashCommands[margs.SubCommand]
		slashSubCommands := []*discordgo.ApplicationCommandOption{
			&CurrentWeatherApplicationData,
			&CurrentTemperaturApplicationData,
			&CurrentWindspeedApplicationData,
			&APODApplicationData,
			&SeaIceApplicationData,
			&SeaLevelApplicationData,
			&GlobalTempApplicationData,
			&CO2ApplicationData,
		}
		slashData, f := findSubSlashData(slashSubCommands, margs.SubCommand)
		if !f {
			err := s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Unable to find that sub-command name, please report this with **`Problem`** button!", nil, nil))
			if err != nil {
				utils.HandleClientError(s, i, err, "help")
				return
			}
			return
		}

		_, found := utils.CommandBlock.Get(margs.SubCommand)

		embed = []*discordgo.MessageEmbed{
			{
				Color:       0xf2c56b,
				Title:       CapitalizeTitle(slashData.Name),
				Description: slashData.Description + fmt.Sprintf("\n\n**Arguments**\n%v", ArgumentToString(slashData.Options)),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Slowdown",
						Value:  fmt.Sprintf("%v ms", commandData.Ratelimit),
						Inline: true,
					},
					{
						Name:   "Disabled",
						Value:  fmt.Sprintf("%v", found),
						Inline: true,
					},
				},
			},
		}
	} else {
		embed = []*discordgo.MessageEmbed{
			{
				Color:       0xf2c56b,
				Title:       "Geobot",
				Description: fmt.Sprintf("A educational discord bot, it allows user to learn geography, astronomical and geometric. Utilizes many APIs from NASA, Google and Weather.\n\nSelect a **command** or a **sub-command** to use get more information about that command!"),
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "From Geobot team",
					IconURL: s.State.User.AvatarURL(""),
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: s.State.User.AvatarURL(""),
				},
			},
		}
	}
	if margs.Public == true {
		err := s.InteractionRespond(i.Interaction, utils.SendEmbed(embed, nil))
		if err != nil {
			utils.HandleClientError(s, i, err, "help")
			return
		}
	} else {
		err := s.InteractionRespond(i.Interaction, utils.SendPrivateEmbed(embed, nil))
		if err != nil {
			utils.HandleClientError(s, i, err, "help")
			return
		}
	}
}

func findSlashData(name string) (*discordgo.ApplicationCommand, bool) {
	var res *discordgo.ApplicationCommand

	for _, v := range SlashCommands() {
		if v.Name == name {
			res = v
			return res, true
		}
	}
	return nil, false
}

func findSubSlashData(data []*discordgo.ApplicationCommandOption, name string) (*discordgo.ApplicationCommandOption, bool) {
	var res *discordgo.ApplicationCommandOption

	for _, v := range data {
		if v.Name == name {
			res = v
			return res, true
		}
	}
	return nil, false
}

func ArgumentToString(arr []*discordgo.ApplicationCommandOption) string {
	var options []string
	for _, option := range arr {
		if len(option.Choices) > 0 {
			options = append(options, fmt.Sprintf("- **`%v`** %v\n%v\n> %v", option.Name, RequiredValidator(option.Required), option.Description, ChoicesToArrayString(option.Choices)))
		} else {
			options = append(options, fmt.Sprintf("- **`%v`** %v\n%v", option.Name, RequiredValidator(option.Required), option.Description))
		}
	}

	return strings.Join(options, "\n")
}

func ChoicesToArrayString(choices []*discordgo.ApplicationCommandOptionChoice) string {

	arr := make([]string, len(choices))

	for _, v := range choices {
		arr = append(arr, fmt.Sprintf("`%v`, ", v.Name))
	}

	return strings.Join(arr, "")
}

func RequiredValidator(required bool) string {
	if required {
		return "**Required**"
	} else {
		return "**Not required**"
	}
}

func CapitalizeTitle(s string) string {
	return fmt.Sprintf("%v%v", strings.ToUpper(s[0:1]), s[1:])
}
