package events

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"github.com/bwmarrin/discordgo"
)

var (
	Ratelimit = class.NewRatelimit()
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.SendLogMessage("The bot is running!")
	utils.HandleDebugMessage("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	s.UpdateListeningStatus("slash commands")
}
