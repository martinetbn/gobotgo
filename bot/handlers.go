package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/martinetbn/gobotgo/events"
)

func Handlers(discord *discordgo.Session) {
	discord.AddHandler(events.MessageCreate)
	discord.AddHandler(events.Ready)
	discord.AddHandler(events.InteractionCreate)
}
