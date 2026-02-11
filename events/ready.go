package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/commands/slash"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is connected.")
	slash.RegisterPing(s)
}
