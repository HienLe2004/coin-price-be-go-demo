package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

// Cấu trúc dữ liệu từ Binance
type BinanceMessage struct {
	Price string `json:"p"`
}

// Cấu trúc dữ liệu cho frontend
type CryptoData struct {
	Price float64 `json:"price"`
}

// Các biến toàn cục
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan CryptoData)
var upgrader = websocket.Upgrader{}

func main() {
	// Kết nối WebSocket của Binance
	go connectBinanceWebSocket()

	// Tạo WebSocket cho frontend
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func connectBinanceWebSocket() {
	// Kết nối tới Binance WebSocket cho BTC/USDT
	binanceURL := "wss://stream.binance.com:9443/ws/btcusdt@trade"
	conn, _, err := websocket.DefaultDialer.Dial(binanceURL, nil)
	if err != nil {
		log.Fatal("Could not connect to Binance WebSocket:", err)
	}
	defer conn.Close()

	for {
		// Nhận dữ liệu từ Binance
		var msg BinanceMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading Binance WebSocket message:", err)
			break
		}

		// Chuyển đổi giá từ string sang float64
		price, err := strconv.ParseFloat(msg.Price, 64)
		if err != nil {
			log.Println("Error parsing price:", err)
			continue
		}

		// Gửi giá mới qua kênh broadcast cho frontend
		broadcast <- CryptoData{Price: price}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Tạo WebSocket kết nối với client
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Lưu kết nối vào danh sách client
	clients[ws] = true

	// Đọc dữ liệu từ client để duy trì kết nối
	for {
		var msg CryptoData
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
	}
}

func handleMessages() {
	for {
		// Nhận giá mới từ kênh broadcast
		data := <-broadcast

		// Gửi dữ liệu tới tất cả client đang kết nối
		for client := range clients {
			err := client.WriteJSON(data)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
