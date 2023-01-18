package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.Println("Joined ", g.Name, g.ID, g.OwnerID)

	DBL.PostStats(len(s.State.Guilds))
}
