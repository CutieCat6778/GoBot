package components

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ComponentsHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
)

func init() {
	ComponentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"sealevel_left":  GoLeft,
		"sealevel_right": GoRight,
	}
}
