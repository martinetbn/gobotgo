package slash

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/registry"
)

func init() {
	registry.RegisterSlashCommand("ping", "Replies with pong.", Ping)
}

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
