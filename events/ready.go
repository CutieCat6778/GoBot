package events

import (
	"cutiecat6778/discordbot/api"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"log"
)

var (
	Ratelimit = class.NewRatelimit()
	DBL       api.DBL
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.SendLogMessage("The bot is running!")
	utils.HandleDebugMessage(fmt.Sprintf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator))
	s.UpdateListeningStatus("slash commands")

	DBL = api.NewDBL()
	log.Println("Posted stats ", len(s.State.Guilds))
	DBL.PostStats(len(s.State.Guilds))
	api.ListenVotes()

	c := cron.New()
	_, err := c.AddFunc("@hourly", func() {
		log.Println("Posted stats ", len(s.State.Guilds))
		DBL.PostStats(len(s.State.Guilds))
	})
	if err != nil {
		utils.HandleServerError(err)
	}
	c.Start()
}
