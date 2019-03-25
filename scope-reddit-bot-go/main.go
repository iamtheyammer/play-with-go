package main

import (
  //graw "github.com/turnage/graw"
  reddit "github.com/turnage/graw/reddit"
  "fmt"
  "time"
  "strconv"
)

var cfg reddit.BotConfig = reddit.BotConfig{
    Agent: "iamtheyammer",
    // Your registered app info from following:
    // https://github.com/reddit/reddit/wiki/OAuth2
    App: reddit.App{
        ID:       "sP1UTOiyknBKrQ",
        Secret:   "sWRa9Eyd0xbjgcaccqkkt0x1R0k",
        Username: "iamtheyammer",
        Password: "rti6XNMQUdlpuhol6jAQ3NAR",
    },
    Config: []string{"applyingtocollege"}
}

var bot reddit.Bot

func initBot() {
  initB, err := reddit.NewBot(cfg)
  if err != nil {
    fmt.Println(err)
  }
  bot = initB
}

func main() {
    initBot()
    bot.SendMessage("samtheboo", "Working...", strconv.Itoa(int(time.Now().Unix())))
}
