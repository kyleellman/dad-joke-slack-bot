## Dad Joke Slack Bot
When I learned that fatherhood.gov has an API that returns dad jokes, I decided to created a slack bot that posts dad jokes as an excuse to practice some topics I've been wanting to learn more about:
- Golang basics
- Slack apps
- Github actions

### What it does
The bot queries the fatherhood.gov API, selects a random joke, and posts it to slack. The bot sends the opener in a message, and the punchline in a thread reply to give the viewer a chance to guess.

The bot is not interactive, but posts the joke when a github action runs, which is on a schedule.

### Notes on setup
- This requires a slack app. The token is an oauth token from the slack app.
- The app requires the `chat:write` scope.
- The app must be installed onto the channel it should run in during the Oauth flow.
- The app must also be "added" to the slack channel from the slack app/web.
- The channel ID can be obtained from the last segment of the channel's URL.
