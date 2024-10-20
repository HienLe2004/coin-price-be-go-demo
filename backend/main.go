package main

import (
	//"fmt"
	"log"
	//"net/url"
	"github.com/gorilla/websocket"
)

const (
	wsServerURL = "wss://stream.binance.com:9443/ws/btcusdt@kline_1m"
)

func main() {
	log.Printf("connecting to %s", wsServerURL)

	// Thiết lập kết nối
	c, _, err := websocket.DefaultDialer.Dial(wsServerURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Vòng lặp để nhận dữ liệu
	i := 1
	for i < 10 {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		i++
		// Giải mã dữ liệu và xử lý ở đây
	}
}
