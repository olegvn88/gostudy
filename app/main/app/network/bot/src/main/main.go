package main

import (
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//make sure your Procfile contains "web: main" and located in root app directory
type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Get Joke"},
}

//const WebhookURL = "https://english891.herokuapp.com/"
const s = "Welcome to the help menu of boot english891!\n - get joke \n - time \n - help \n"

func getJoke() string {
	c := http.Client{}
	resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
	if err != nil {
		return "Jokes API not responding"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Joke error"
	}
	return joke.Value.Joke
}

func main() {
	// Heroku прокидывает порт для приложения в переменную окружения PORT
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI("735585631:AAEyopXfFtcaZj3wD59G_JYmQabs-WwNj-g")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//// Устанавливаем вебхук
	//_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	//if err != nil {
	//	log.Fatal(err)
	//}


//===
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
//======
	//updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	//получаем все обновления из канала updates
	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("received text: ", update.Message.Text)

		switch strings.ToLower(update.Message.Text) {
		case "bot":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "hi")
		case "help":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, s)
		case "get joke":
			//если пользователь нажал на кнопку,то прийдет сообщение "Get Joke"
			message = tgbotapi.NewMessage(update.Message.Chat.ID, getJoke())
		case "time":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, time.Now().String())
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `Press Get Joke to receive joke or type help`)
		}

		//в ответном сообщении просим показать клавиатуру

		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		_, _ = bot.Send(message)

	}
}

//https://api.telegram.org/bot735585631:AAEyopXfFtcaZj3wD59G_JYmQabs-WwNj-g/setwebhook?url=https://english891.herokuapp.com/
