package utils

import "github.com/bwmarrin/discordgo"

func SendInteractionMessage(c string, e []*discordgo.MessageEmbed, m *discordgo.MessageAllowedMentions) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         c,
			Embeds:          e,
			AllowedMentions: m,
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
		},
	}
}

func SendEmbed(e []*discordgo.MessageEmbed, m *discordgo.MessageAllowedMentions) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:          e,
			AllowedMentions: m,
		},
	}
}
