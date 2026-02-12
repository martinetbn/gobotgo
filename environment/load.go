package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DeveloperID string
	ClientID    string
	GuildID     string
	Token       string
	Prefix      string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	DeveloperID = os.Getenv("BOT_DEVELOPER_ID")
	ClientID = os.Getenv("BOT_CLIENT_ID")
	GuildID = os.Getenv("BOT_GUILD_ID")
	Token = os.Getenv("BOT_TOKEN")
	Prefix = os.Getenv("BOT_PREFIX")
}
