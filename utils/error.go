package utils

import (
	"cutiecat6778/discordbot/class"
	"log"
	"os"

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
	SendErrorMessage("[Client] [Error] ", err.Error())
	Error.Println(err)
	CommandBlock.Register(name, true)

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
		Error.Println(sent)
	}
}

func HandleClientBlock(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	SendErrorMessage("[Server] [Error]", err.Error())
	Error.Fatal(err)
}

func HandleInfoMessage(msg ...any) {
	Info.Println(msg)
}

func HandleDebugMessage(msg ...any) {
	Debug.Println(msg)
}

func HandleWarningMessage(msg ...any) {
	Warning.Println(msg)
}
