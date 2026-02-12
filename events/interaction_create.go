package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/registry"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	registry.ExecuteSlashCommand(s, i)
}
