package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/i3odja/telegram-bot/chatbot"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const debugMode = true

var baseURL = "https://cryptic-brushlands-39510.herokuapp.com/"

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	// to monitor changes run: heroku logs --tail
	log.Printf("From: %+v Text: %+vn", update.Message.From, update.Message.Text)
}

func initTelegram() {
	os.Setenv("TOKEN_TG_BOT", "1101236908:AAGdRKCvt8EzpByAFjPKnof-gYKjdTE9jVM")
	bot, err := chatbot.CreateNewBotConnection()
	if err != nil {
		log.Fatal("cannot connect to bot %w", err)
	}

	// this perhaps should be conditional on GetWebhookInfo()
	// only set webhook if it is not set properly
	url := baseURL + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8099"
	}

	// gin router
	router := gin.New()
	router.Use(gin.Logger())

	// telegram
	initTelegram()
	token := os.Getenv("TOKEN_TG_BOT")
	router.POST("/" + token, webhookHandler)

	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}

//func AllCommands(w http.ResponseWriter, r *http.Request) {
//	chatbot.CommandList()
//}


//func Start() {
//	r := mux.NewRouter()
//	r.HandleFunc("/help", AllCommands)
//	r.HandleFunc("/products", ProductsHandler)
//	r.HandleFunc("/articles", ArticlesHandler)
//	http.Handle("/", r)
//}
