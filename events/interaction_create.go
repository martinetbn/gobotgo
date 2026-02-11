package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/commands/slash"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.Ping(s, i)
}
