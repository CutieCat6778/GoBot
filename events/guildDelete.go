package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.Println("Left ", g.Name, g.ID, g.OwnerID)

	DBL.PostStats(len(s.State.Guilds))
}
