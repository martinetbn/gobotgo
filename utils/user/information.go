package user

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GetID(u *discordgo.User) string {
	return u.ID
}

func GetDisplayName(u *discordgo.User) string {
	return u.DisplayName()
}

func GetUsername(u *discordgo.User) string {
	return u.String()
}

func GetColor(s *discordgo.Session, u *discordgo.User) int {
	fullUser, err := s.User(u.ID)
	if err != nil {
		return 0
	}
	return fullUser.AccentColor
}

func GetCreatedAt(u *discordgo.User) time.Time {
	id, _ := strconv.ParseInt(u.ID, 10, 64)
	ms := (id >> 22) + 1420070400000
	return time.UnixMilli(ms)
}

func GetJoinedAt(s *discordgo.Session, guildID string, u *discordgo.User) time.Time {
	member, err := s.GuildMember(guildID, u.ID)
	if err != nil {
		return time.Time{}
	}
	return member.JoinedAt
}
