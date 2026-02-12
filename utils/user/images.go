package user

import "github.com/bwmarrin/discordgo"

func GetAvatarUrl(u *discordgo.User) string {
	return u.AvatarURL("1024")
}

func GetBannerUrl(s *discordgo.Session, u *discordgo.User) string {
	fullUser, err := s.User(u.ID)
	if err != nil {
		return ""
	}
	return fullUser.BannerURL("4096")
}
