package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/commands/registry"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is connected.")
	registry.PushRegisteredSlashCommands(s)
}
