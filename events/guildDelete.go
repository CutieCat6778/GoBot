package events

import (
	"cutiecat6778/discordbot/class"
	"github.com/bwmarrin/discordgo"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildCreate) {
	utils.Debug.Println("Left ", g.Name, g.ID, g.OwnerID)

	if !class.Ignore {
		DBL.PostStats(len(s.State.Guilds))
	}
}
