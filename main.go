package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kh9543/koala/apps/bot"
	"github.com/kh9543/koala/apps/http/heartbeat"
	"github.com/kh9543/koala/domain/kv/memory"
)

var (
	port     string
	botToken string
)

func init() {
	botToken = os.Getenv("token")
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
}

func main() {
	kv := memory.NewMemory()
	if err := bot.NewDiscordBot("!", botToken, kv); err != nil {
		panic(err)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	heartbeat.NewHandler(apiv1)
	r.Run(fmt.Sprintf(":%s", port))
}
