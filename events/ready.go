package events

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	Ratelimit = class.NewRatelimit()
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.SendLogMessage("The bot is running!")
	utils.HandleDebugMessage(fmt.Sprintf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator))
	s.UpdateListeningStatus("slash commands")
}
