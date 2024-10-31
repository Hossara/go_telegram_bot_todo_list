package client

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
	"log"
	"net"
	"net/http"
	"os"
)

func Init() *tgbotapi.BotAPI {
	// Set up SOCKS5 proxy
	if os.Getenv("PROXY") != "" {
		// Create socks 5 connection
		dialer, err := proxy.SOCKS5("tcp", os.Getenv("PROXY"), &proxy.Auth{
			User:     "",
			Password: "",
		}, proxy.Direct)

		if err != nil {
			log.Fatalf("failed to connect to proxy: %v", err)
			return nil
		}

		// Use it in http transport layer
		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}

		// Create HTTP client with the proxy
		client := &http.Client{Transport: transport}
		bot, err := tgbotapi.NewBotAPIWithClient(os.Getenv("TELEGRAM_APITOKEN"), client)

		if err != nil {
			log.Panic(err)
			return nil
		}

		return bot
	} else {
		// Create telegram bot instance without proxy
		bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))

		if err != nil {
			log.Panic(err)
			return nil
		}

		return bot
	}
}
