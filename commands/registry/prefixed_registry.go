package registry

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/environment"
)

type prefixedCommand struct {
	Name        string
	Description string
	Handler     func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

var prefixedCommandList = []prefixedCommand{}

func RegisterPrefixedCommand(name, description string, handler func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {
	prefixedCommandList = append(prefixedCommandList, prefixedCommand{
		Name:        name,
		Description: description,
		Handler:     handler,
	})
}

func ExecutePrefixedCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, command := range prefixedCommandList {
		if strings.HasPrefix(m.Content, environment.Prefix+command.Name) {
			command.Handler(s, m, strings.Fields(m.Content[len(environment.Prefix+command.Name):]))
			return
		}
	}
}
