package events

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	Ratelimit = class.NewRatelimit()
	DBL       api.DBL
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.SendLogMessage("The bot is running!")
	utils.HandleDebugMessage(fmt.Sprintf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator))
	s.UpdateListeningStatus("slash commands")

	for _, v := range r.Guilds {
		utils.Debug.Println(v.Name, v.JoinedAt, v.MemberCount)
	}

	if class.LOCAL {
		return
	}

	DBL = api.NewDBL()
	utils.Debug.Println("Posted stats ", len(s.State.Guilds))
	DBL.PostStats(len(s.State.Guilds))

	api.ListenVotes()
}
