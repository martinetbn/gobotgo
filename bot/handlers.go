package bot

import (
	"github.com/bwmarrin/discordgo"
	_ "github.com/martinetbn/gobotgo/commands/prefixed"
	_ "github.com/martinetbn/gobotgo/commands/slash"
	"github.com/martinetbn/gobotgo/events"
)

func Handlers(discord *discordgo.Session) {
	discord.AddHandler(events.MessageCreate)
	discord.AddHandler(events.Ready)
	discord.AddHandler(events.InteractionCreate)
}
