package main

import (
	"github.com/martinetbn/gobotgo/bot"
	"github.com/martinetbn/gobotgo/environment"
)

func main() {
	environment.Load()
	bot.Connect()
}
