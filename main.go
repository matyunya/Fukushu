package main

import (
  "fukushu/Entries"
  "gopkg.in/telegram-bot-api.v4"
  "log"
  "math/rand"
  "os"
  "time"
)

var (
  chatIDs  = make(map[int64]struct{})
  bot, err = tgbotapi.NewBotAPI(os.Getenv("FUKUSHU_TOKEN"))
)

// todo: save chatIDs
func getUpdates() {
  u := tgbotapi.NewUpdate(0)
  u.Timeout = 5

  updates, _ := bot.GetUpdatesChan(u)
  for update := range updates {
    if update.Message == nil {
      continue
    }

    chatIDs[update.Message.Chat.ID] = struct{}{}

    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    entry := getEntry()
    send(update.Message.Chat.ID, entry.ToString())
  }
}

func send(id int64, message string) {
  msg := tgbotapi.NewMessage(id, message)
  msg.ParseMode = "HTML"
  bot.Send(msg)
}

func getEntry() fukushu.Entry {
  return fukushu.Entries[rand.Intn(len(fukushu.Entries))]
}

func main() {
  rand.Seed(time.Now().Unix())

  if err != nil {
    log.Printf("Error: %v", err)
  }
  log.Printf("Authorized on account %s", bot.Self.UserName)

  go getUpdates()

  for true {
    time.Sleep(1 * time.Minute)
    for ID := range chatIDs {
      entry := getEntry()
      send(ID, entry.ToString())
    }
  }

}
