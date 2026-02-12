package registry

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type SlashCommand struct {
	Name        string
	Description string
	Handler     func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var slash_command_list = []SlashCommand{}

func RegisterSlashCommand(name, description string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	slash_command_list = append(slash_command_list, SlashCommand{Name: name, Description: description, Handler: handler})
}

func PushRegisteredSlashCommands(s *discordgo.Session) {
	for _, command := range slash_command_list {
		s.ApplicationCommandCreate(s.State.User.ID, os.Getenv("GUILD_ID"), &discordgo.ApplicationCommand{
			Name:        command.Name,
			Description: command.Description,
		})
	}
}

func ExecuteSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, command := range slash_command_list {
		if command.Name == i.ApplicationCommandData().Name {
			command.Handler(s, i)
			return
		}
	}
}
