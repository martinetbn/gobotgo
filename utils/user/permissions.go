package user

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/environment"
)

func HasPermission(s *discordgo.Session, m *discordgo.MessageCreate, perm int64) bool {
	perms, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		return false
	}
	return perms&perm != 0
}

func IsDev(m *discordgo.MessageCreate) bool {
	return m.Author.ID == environment.DeveloperID
}
