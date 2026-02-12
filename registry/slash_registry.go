package registry

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/environment"
)

type slashCommand struct {
	Name        string
	Description string
	Handler     func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var slashCommandList = []slashCommand{}

func RegisterSlashCommand(name, description string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	slashCommandList = append(slashCommandList, slashCommand{Name: name, Description: description, Handler: handler})
}

func PushRegisteredSlashCommands(s *discordgo.Session) {
	for _, command := range slashCommandList {
		s.ApplicationCommandCreate(s.State.User.ID, environment.GuildID, &discordgo.ApplicationCommand{
			Name:        command.Name,
			Description: command.Description,
		})
	}
}

func ExecuteSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, command := range slashCommandList {
		if command.Name == i.ApplicationCommandData().Name {
			command.Handler(s, i)
			return
		}
	}
}
