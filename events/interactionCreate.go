package events

import (
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/components"
	"cutiecat6778/discordbot/database"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		InteractionApplicationCommand(s, i)
	case discordgo.InteractionMessageComponent:
		InteractionMessageComponent(s, i)
	}
}

func InteractionApplicationCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slashHandlers := commands.SlashHandlers()
	g, f := database.FindServerByServerID(i.GuildID)
	if !f {
		id := database.CreateGuild(i.GuildID)
		g, _ = database.FindServerByID(id)
	}

	name := i.ApplicationCommandData().Name

	if h, ok := slashHandlers[i.ApplicationCommandData().Name]; ok {
		// Ratelimit
		current_time := time.Now().Unix()
		r, f := Ratelimit.Get(i.Member.User.ID)
		if !f {
			Ratelimit.Register(i.Member.User.ID)
			r, f = Ratelimit.Get(i.Member.User.ID)
		}
		time := r.GetTime()

		if len(i.ApplicationCommandData().Options) > 1 && len(i.ApplicationCommandData().Options[0].Name) > 0 {
			h.Data = h.Data.SubCommandData[i.ApplicationCommandData().Options[0].Name]
			name = i.ApplicationCommandData().Options[0].Name
		}

		c, f1 := utils.CommandBlock.Get(name)
		if f1 {
			status := c.GetStatus()
			if status {
				utils.HandleClientBlock(s, i)
				return
			}
		}

		if current_time-time >= h.Data.Ratelimit && f {

			// Execute command
			h.Execute(s, i, g)
			Ratelimit.Write(i.Member.User.ID)
		} else {
			s.InteractionRespond(i.Interaction, utils.SendPrivateInteractionMessage("Please slow down, don't use this command to fast! Please wait "+fmt.Sprint((current_time-time)-h.Data.Ratelimit)+" seconds", nil, nil))
			return
		}
	}
}

func InteractionMessageComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.MessageComponentData().CustomID

	if strings.HasPrefix(name, "sealevel") {
		name = name[0:8]
	}

	utils.HandleDebugMessage(name)

	if h, ok := components.ComponentsHandlers[name]; ok {
		h(s, i)
	}
}
