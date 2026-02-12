package registry

import (
	"os"

	"github.com/bwmarrin/discordgo"
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
		if m.Content == os.Getenv("BOT_PREFIX")+command.Name {
			command.Handler(s, m)
			return
		}
	}
}
