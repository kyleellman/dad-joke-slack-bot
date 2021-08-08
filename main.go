package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/kyleellman/dad-joke-slack-bot/pkg/jokes"
	"github.com/kyleellman/dad-joke-slack-bot/pkg/slack"
)

func main() {
	// Seed random to ensure a different joke is selected each time.
	rand.Seed(time.Now().UnixNano())

	// Fetch a random one.
	joke := jokes.GetRandomJoke()

	// TODO use token/channel as env vars
	ssc := slack.NewSimpleClient(os.Getenv("TOKEN"), os.Getenv("CHANNEL_ID"))

	// Send the joke opener and retain the ersult for the thread_ts
	smr := ssc.SendMessage(joke.Opener, nil)

	// Send the joke punchline as a reply to the original slack thread
	ssc.SendMessage(joke.Punchline, &smr.Ts)
}
