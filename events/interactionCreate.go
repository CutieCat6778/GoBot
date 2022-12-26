package events

import (
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slashHandlers := commands.SlashHandlers()
	g, f := database.FindServerByServerID(i.GuildID)
	if !f {
		id := database.CreateGuild(i.GuildID)
		g, _ = database.FindServerByID(id)
	}

	if h, ok := slashHandlers[i.ApplicationCommandData().Name]; ok {
		// Ratelimit
		current_time := time.Now().Unix()
		r, f := Ratelimit.Get(i.Member.User.ID)
		if !f {
			Ratelimit.Register(i.Member.User.ID)
			r, f = Ratelimit.Get(i.Member.User.ID)
		}
		time := r.GetTime()

		if len(i.ApplicationCommandData().Options[0].Name) > 0 {
			h.Data = h.Data.SubCommandData[i.ApplicationCommandData().Options[0].Name]
		}

		if current_time-time >= h.Data.Ratelimit && f {

			// Execute command
			h.Execute(s, i, g)
		} else {
			s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Please slow down, don't use this command to fast! Please wait "+fmt.Sprint((current_time-time)-h.Data.Ratelimit)+" seconds", nil, nil))
			return
		}
	}
}
