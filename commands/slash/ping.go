package slash

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "ping" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	}
}

func RegisterPing(s *discordgo.Session) {
	s.ApplicationCommandCreate(s.State.User.ID, os.Getenv("GUILD_ID"), &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with Pong!",
	})
}
