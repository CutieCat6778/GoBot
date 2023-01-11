package components

import "github.com/bwmarrin/discordgo"

func Error(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "error_survey_" + i.Interaction.Member.User.ID,
			Title:    "Error questions",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "command_name",
							Label:       "What is the command name?",
							Style:       discordgo.TextInputShort,
							Placeholder: "Command name (e.g. /ping, ping)",
							Required:    true,
							MaxLength:   30,
							MinLength:   4,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:  "problem",
							Label:     "Description about the problem/suggestion?",
							Style:     discordgo.TextInputParagraph,
							Required:  true,
							MaxLength: 2000,
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
