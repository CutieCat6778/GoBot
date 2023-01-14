package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	EarthApplicationData = discordgo.ApplicationCommandOption{
		Name:        "earth",
		Description: "Astronomy Picture of the Day",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "year",
				Description: "What year you want earth image to be?",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "month",
				Description: "What month you want earth image to be?",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "day",
				Description: "What day you want earth image to be?",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "private",
				Description: "Should it display only for you or public?",
				Required:    false,
			},
		},
		Type: discordgo.ApplicationCommandOptionSubCommand,
	}
	EarthCommandData class.CommandData
)

type EarthOption struct {
	Private bool
	Year    string
	Month   string
	Day     string
}

func init() {
	var (
		defaultPerms int64 = discordgo.PermissionSendMessages
	)

	EarthCommandData = class.CommandData{
		Permissions: defaultPerms,
		Ratelimit:   5000,
		BotPerms:    defaultPerms,
	}
}

func Earth(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {
	options := i.ApplicationCommandData().Options[0].Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := EarthOption{}
	if option, ok := optionMap["year"]; ok {
		margs.Year = option.StringValue()
	}
	if option, ok := optionMap["month"]; ok {
		margs.Month = option.StringValue()
	}
	if option, ok := optionMap["day"]; ok {
		margs.Day = option.StringValue()
	}
	if option, ok := optionMap["private"]; ok {
		margs.Private = option.BoolValue()
	}

	s.InteractionRespond(i.Interaction, utils.DeferInteraction())

	trigger := false
	var date string
	if len(margs.Year)+len(margs.Month)+len(margs.Day) < 6 && len(margs.Year)+len(margs.Month)+len(margs.Day) > 0 {
		_, err := s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: "You need to enter a **full date**, with 3 arguments `year`, `month` and `date`!!",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		})

		if err != nil {
			utils.HandleClientError(s, i, err, "earth")
		}
		return
	} else if len(margs.Year)+len(margs.Month)+len(margs.Day) == 0 {
		date = time.Now().Format("2006-01-02")
	} else {
		if len(margs.Month) == 1 {
			margs.Month = "0" + margs.Month
		}
		if len(margs.Day) == 1 {
			margs.Day = "0" + margs.Day
		}
		date = fmt.Sprintf("%v-%v-%v", margs.Year, margs.Month, margs.Day)
	}

	log.Println(date)

	data := AstronomyClass.Earth(date)
	if len(data) == 0 {
		data = AstronomyClass.Earth2()
		trigger = true
	}
	if len(data) == 0 {
		utils.HandleClientError(s, i, errors.New("error while fetching earth image"), "earth")
		return
	}
	url := AstronomyClass.EarthImage(data[0].Date, data[0].Image)
	height, width := AstronomyClass.GetImageSize(url)

	log.Println(height, width)

	res := []*discordgo.MessageEmbed{
		{
			Title:       "Earth image",
			Description: data[0].Caption,
			Color:       0xf2c56b,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Was taken in " + strings.Split(data[0].Date, " ")[0],
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    url,
				Width:  width,
				Height: height,
			},
		},
	}
	if margs.Private {
		_, err := s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Embeds: res,
			Flags:  discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		})

		if err != nil {
			utils.HandleClientError(s, i, err, "earth")
		}
	} else {
		_, err := s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Embeds: res,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		})
		if err != nil {
			utils.HandleClientError(s, i, err, "earth")
		}
	}
	if trigger {
		time.Sleep(time.Second)
		_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "There was no image for that day, so I showed you the most recent image!",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		})
		if err != nil {
			utils.HandleClientError(s, i, err, "earth")
			return
		}
	}
}
