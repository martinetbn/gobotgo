package registry

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/environment"
)

type PrefixedCommand struct {
	Name        string
	Description string
	Handler     func(s *discordgo.Session, m *discordgo.MessageCreate)
}

var prefixed_command_list = []PrefixedCommand{}

func RegisterPrefixedCommand(name, description string, handler func(s *discordgo.Session, m *discordgo.MessageCreate)) {
	prefixed_command_list = append(prefixed_command_list, PrefixedCommand{
		Name:        name,
		Description: description,
		Handler:     handler,
	})
}

func ExecutePrefixedCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, command := range prefixed_command_list {
		if m.Content == environment.Prefix+command.Name {
			command.Handler(s, m)
			return
		}
	}
}
