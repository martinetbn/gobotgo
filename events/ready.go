package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/commands/registry"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("Bot is ready.")
	registry.PushRegisteredSlashCommands(s)
}
