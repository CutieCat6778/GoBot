package modal

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ModalHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
)

func init() {
	ModalHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"error": ErrorHandler,
	}
}
