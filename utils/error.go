package utils

import (
	"cutiecat6778/discordbot/class"
	"log"
	"os"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

var (
	Info         = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
	Warning      = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
	Error        = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	Debug        = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
	CommandBlock = class.NewCommandBlock()
)

func HandleClientError(s *discordgo.Session, i *discordgo.InteractionCreate, err error, name string) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		Debug.Printf("called from %s#%d\n", file, no)
	}
	SendErrorMessage("<@"+class.OwnerID+">", err.Error())
	Error.Println(err)
	if len(name) > 0 {
		CommandBlock.Register(name, true)
	}

	embed := []*discordgo.MessageEmbed{
		{
			Color:       0xf2c56b,
			Description: "**502 | Internal Server Error**\n\nSorry for the inconvenient, the developer is now informed about this problem. There will be a fix soon!",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "From GeoBot developer ",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    s.State.User.AvatarURL(""),
				Width:  200,
				Height: 200,
			},
		},
	}

	sent := s.InteractionRespond(i.Interaction, SendPrivateEmbed(embed, nil))
	if sent != nil {
		_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Embeds: embed,
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
			Error.Println(err)
		}
	}
}

func HandleClientBlock(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		Debug.Printf("called from %s#%d\n", file, no)
	}
	embed := []*discordgo.MessageEmbed{
		{
			Color:       0xf2c56b,
			Description: "**Command under construction**\n\nSorry for the inconvenient, there is a problem with this command currently. The developer has been informed, please try again later!",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "From GeoBot developer ",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    s.State.User.AvatarURL(""),
				Width:  200,
				Height: 200,
			},
		},
	}

	sent := s.InteractionRespond(i.Interaction, SendPrivateEmbed(embed, nil))

	if sent != nil {
		Error.Println(sent)
	}
}

func HandleServerError(err error) {
	SendErrorMessage("<@"+class.OwnerID+">", err.Error())
	_, file, no, ok := runtime.Caller(1)
	if ok {
		Debug.Printf("called from %s#%d\n", file, no)
	}
	Error.Println(err)
}

func HandleInfoMessage(msg ...any) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		Debug.Printf("called from %s#%d\n", file, no)
	}
	Info.Println(msg)
}

func HandleDebugMessage(msg ...any) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		Debug.Printf("called from %s#%d\n", file, no)
	}
	Debug.Println(msg)
}

func HandleWarningMessage(msg ...any) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		Debug.Printf("called from %s#%d\n", file, no)
	}
	Warning.Println(msg)
}
