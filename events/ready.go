package events

import (
	"cutiecat6778/discordbot/utils"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Fatal("The bot is ready!")
	utils.SendLogMessage("Bot is ready!")
}
