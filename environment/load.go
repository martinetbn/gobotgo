package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ClientID string
	GuildID  string
	Token    string
	Prefix   string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	ClientID = os.Getenv("CLIENT_ID")
	GuildID = os.Getenv("GUILD_ID")
	Token = os.Getenv("DISCORD_TOKEN")
	Prefix = os.Getenv("BOT_PREFIX")
}
