package utils

import "github.com/bwmarrin/discordgo"

func SendInteractionMessage(c string, e []*discordgo.MessageEmbed, m *discordgo.MessageAllowedMentions) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         c,
			Embeds:          e,
			AllowedMentions: m,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		},
	}
}

func SendPrivateInteractionMessage(c string, e []*discordgo.MessageEmbed, m *discordgo.MessageAllowedMentions) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         c,
			Embeds:          e,
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: m,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		},
	}
}

func SendPrivateEmbed(e []*discordgo.MessageEmbed, m *discordgo.MessageAllowedMentions) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:          e,
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: m,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		},
	}
}

func SendEmbed(e []*discordgo.MessageEmbed, m *discordgo.MessageAllowedMentions) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:          e,
			AllowedMentions: m,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		},
	}
}

func DeferInteraction() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Loading...",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Problem?",
							Style:    discordgo.SecondaryButton,
							CustomID: "error",
						},
					},
				},
			},
		},
	}
}
