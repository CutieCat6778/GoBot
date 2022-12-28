package components

import (
	"cutiecat6778/discordbot/commands"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func GoLeft(s *discordgo.Session, i *discordgo.InteractionCreate) {

	user, found := commands.SeaLevelScroll.Get(i.Member.User.ID)
	if !found {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Description: "There is a problem, please report this problem!",
				},
			},
		})
		return
	}

	position := i.MessageComponentData().CustomID
	position = position[12:]

	num, _ := strconv.ParseInt(position, 6, 12)

	url := commands.URLResolver(user.Location, num+1)
	height, width := commands.AstronomyClass.GetImageSize(url)

	embed := []*discordgo.MessageEmbed{
		{
			Title:       "Sea level prediction " + position,
			Color:       0xf2c56b,
			Description: "Recent satellite observations have detected that the Greenland and Antarctic ice sheets are losing ice. Even a partial loss of these ice sheets would cause a 1-meter (3-foot) rise. If lost completely, both ice sheets contain enough water to raise sea level by 66 meters (217 feet).\n\nThis visualization shows the effect on coastal regions for each meter of sea level rise, up to 6 meters (19.7 feet). Land that would be covered in water is shaded red.\n\n[Resources](https://climate.nasa.gov/interactives/climate-time-machine)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Center for Remote Sensing of Ice Sheets",
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    url,
				Width:  width,
				Height: height,
			},
		},
	}

	var component []discordgo.MessageComponent
	if num-1 <= 0 {
		component = []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "←",
				Style:    discordgo.PrimaryButton,
				Disabled: true,
				CustomID: "sealevel_left" + fmt.Sprintf("%v", num),
			},
			discordgo.Button{
				Label:    "→",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "sealevel_right" + fmt.Sprintf("%v", num),
			},
		}
	} else if num-1 > 0 {
		component = []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "←",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "sealevel_left" + fmt.Sprintf("%v", num-1),
			},
			discordgo.Button{
				Label:    "→",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "sealevel_right" + fmt.Sprintf("%v", num-1),
			},
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: component,
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}
}

func GoRight(s *discordgo.Session, i *discordgo.InteractionCreate) {

	user, found := commands.SeaLevelScroll.Get(i.Member.User.ID)
	if !found {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Description: "There is a problem, please report this problem!",
				},
			},
		})
	}
	position := i.MessageComponentData().CustomID
	position = position[14:]
	num, _ := strconv.ParseInt(position, 6, 12)

	url := commands.URLResolver(user.Location, num+1)
	height, width := commands.AstronomyClass.GetImageSize(url)
	embed := []*discordgo.MessageEmbed{
		{
			Title:       "Sea level prediction",
			Color:       0xf2c56b,
			Description: "Recent satellite observations have detected that the Greenland and Antarctic ice sheets are losing ice. Even a partial loss of these ice sheets would cause a 1-meter (3-foot) rise. If lost completely, both ice sheets contain enough water to raise sea level by 66 meters (217 feet).\n\nThis visualization shows the effect on coastal regions for each meter of sea level rise, up to 6 meters (19.7 feet). Land that would be covered in water is shaded red.\n\n[Resources](https://climate.nasa.gov/interactives/climate-time-machine)",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Center for Remote Sensing of Ice Sheets",
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    url,
				Width:  width,
				Height: height,
			},
		},
	}

	var component []discordgo.MessageComponent

	if num+1 >= 6 {
		component = []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "←",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "sealevel_left" + fmt.Sprintf("%v", num),
			},
			discordgo.Button{
				Label:    "→",
				Style:    discordgo.PrimaryButton,
				Disabled: true,
				CustomID: "sealevel_right" + fmt.Sprintf("%v", num),
			},
		}
	} else if num+1 < 6 {
		component = []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "←",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "sealevel_left" + fmt.Sprintf("%v", num+1),
			},
			discordgo.Button{
				Label:    "→",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "sealevel_right" + fmt.Sprintf("%v", num+1),
			},
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: component,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
