package events

import (
	"cutiecat6778/discordbot/class"
	"github.com/bwmarrin/discordgo"
)

func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	utils.Debug.Println("Joined ", g.Name, g.ID, g.OwnerID)

	if !class.Ignore {
		DBL.PostStats(len(s.State.Guilds))
	}
}
