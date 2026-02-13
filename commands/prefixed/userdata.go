package prefixed

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/registry"
	"github.com/martinetbn/gobotgo/utils/user"
)

func init() {
	registry.RegisterPrefixedCommand("userdata", "Show user data.", UserData)
}

func UserData(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: user.GetDisplayName(m.Author) + " Data",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "ID",
				Value:  user.GetID(m.Author),
				Inline: true,
			},
			{
				Name:   "Username",
				Value:  user.GetUsername(m.Author),
				Inline: true,
			},
			{
				Name:  "Account Created",
				Value: user.GetCreatedAt(m.Author).Format(time.DateTime),
			},
			{
				Name:  "Joined Server",
				Value: user.GetJoinedAt(s, m.GuildID, m.Author).Format(time.DateTime),
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     user.GetColor(s, m.Author),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.GetAvatarUrl(m.Author),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: user.GetBannerUrl(s, m.Author),
		},
	})
}
