package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/vsaien/imgo/config"
	"github.com/vsaien/imgo/handler"
	"github.com/vsaien/imgo/models"
)

func main() {
	SafeUsers := new(sync.Map)
	SafeGroup := new(sync.Map)
	c := WebSocketLoadConfig()
	socketMessageHandler := handler.NewWebSocketHandler(
		&models.UserChannelView{
			SafeUsers: SafeUsers,
		},
		&models.DefaultGroup{
			SafeGroup: SafeGroup,
		},
		c,
		func(r *http.Request) bool {
			return true
		},
	)
	log.Println("WebSocket Starting Listen on ..." + c.WebSocketAddress)
	// go socketMessageHandler.Start()
	http.HandleFunc(c.WsPath, socketMessageHandler.Handler)
	http.ListenAndServe(c.WebSocketAddress, nil)
}

func WebSocketLoadConfig() *config.Config {
	configFile := flag.String("f", "config/conf.json", "the config file")
	flag.Parse()
	c := new(config.Config)
	json.Unmarshal([]byte(*configFile), c)
	return c
}
