package events

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 21515ef7e0f18addb1aff961af9152ad15895879
	"github.com/bwmarrin/discordgo"
)

var (
	Ratelimit = class.NewRatelimit()
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.SendLogMessage("The bot is running!")
<<<<<<< HEAD
	utils.HandleDebugMessage("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
=======
	utils.HandleDebugMessage(fmt.Sprintf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator))
>>>>>>> 21515ef7e0f18addb1aff961af9152ad15895879
	s.UpdateListeningStatus("slash commands")
}
