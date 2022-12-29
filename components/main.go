package components

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ComponentsHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
)

func init() {
	ComponentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"sealevel": Move,
	}
}
