package main

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/client/marketwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"time"
)

func main() {

	client := new(client.MarketClient).Init(config.Host)
	optionalRequest := market.GetCandlestickOptionalRequest{Period: market.DAY1, Size: 100}
	resp, err := client.GetCandlestick("btcusdt", optionalRequest)


	client := new(marketwebsocketclient.Last24hCandlestickWebSocketClient).Init(config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func() {
			//client.Request("btcusdt", "1608")
			client.Subscribe("btcusdt", "1608")
		},
		// Response handler
		func(resp interface{}) {
			candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
			if ok {
				if &candlestickResponse != nil {
					if candlestickResponse.Tick != nil {
						t := candlestickResponse.Tick
						applogger.Info("Candlestick update, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}

					if candlestickResponse.Data != nil {
						t := candlestickResponse.Data
						applogger.Info("Candlestick data, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}
				}
			} else {
				applogger.Warn("Unknown response: %v", resp)
			}
		})
	// Connect to the server and wait for the handler to handle the response
	client.Connect(true)
	time.Sleep(time.Minute)
}