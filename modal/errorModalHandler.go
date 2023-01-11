package modal

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func ErrorHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Thank you for taking your time to fill this survey",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		utils.HandleClientError(s, i, err, "")
	}
	data := i.ModalSubmitData()

	if !strings.HasPrefix(data.CustomID, "error_survey_") {
		return
	}

	userid := strings.Split(data.CustomID, "_")[2]
	utils.SendMessage(fmt.Sprintf(
		"<@%s>\n\nFeedback received. From <@%s>\n\n**Command name**:\n%s\n\n**Problem/Suggestion**:\n%s",
		class.OwnerID,
		userid,
		data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
	))
}
