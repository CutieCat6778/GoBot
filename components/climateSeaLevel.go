package components

import (
	"cutiecat6778/discordbot/commands"
	"cutiecat6778/discordbot/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Move(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user, found := commands.SeaLevelScroll.Get(i.Member.User.ID)
	if !found {
		utils.HandleClientError(s, i, errors.New("failed to get user scroll data"), "sealevel")
		return
	}

	position := i.MessageComponentData().CustomID
	move := strings.HasPrefix(position, "sealevel_right")
	if move {
		position = position[14:]
	} else {
		position = position[13:]
	}
	num, _ := strconv.ParseInt(position, 10, 32)
	left, right := ValidateComponent(move, num)
	if move {
		if num+1 < 6 {
			num += 1
		}
	} else {
		if num-1 >= 0 {
			num -= 1
		}
	}
	log.Println(num, move, i.MessageComponentData().CustomID)
	url := commands.URLResolver(user.Location, num)
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

	component := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "←",
			Style:    discordgo.PrimaryButton,
			Disabled: left,
			CustomID: "sealevel_left" + fmt.Sprintf("%v", num),
		},
		discordgo.Button{
			Label:    "→",
			Style:    discordgo.PrimaryButton,
			Disabled: right,
			CustomID: "sealevel_right" + fmt.Sprintf("%v", num),
		},
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
		utils.HandleClientError(s, i, err, "sealevel")
	}
}

func ValidateComponent(move bool, num int64) (bool, bool) {
	if move {
		if num+1 < 6 {
			return false, false
		} else {
			return false, true
		}
	} else {
		if num-1 > 0 {
			return false, false
		} else {
			return true, false
		}
	}
}
