package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/commands"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	commands.Ping(s, m)
}
