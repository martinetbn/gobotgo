package prefixed

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/registry"
)

func init() {
	registry.RegisterPrefixedCommand("echo", "Repeats after the user", Echo)
}

func Echo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Please provide some text to echo.")
		return
	}

	s.ChannelMessageSend(m.ChannelID, strings.Join(args, " "))
}
