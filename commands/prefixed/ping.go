package prefixed

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/commands/registry"
)

func init() {
	registry.RegisterPrefixedCommand("ping", "Replies with Pong.", Ping)
}

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}
