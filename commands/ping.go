package commands

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	PingApplicationData = discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "It will pong you back :)",
	}
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate, g class.Guilds) {

	fmt.Println(g.ServerID)

	s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Pong!", nil, nil))
}
