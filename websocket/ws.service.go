package websocket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	BinanceWS()

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		res := time.Now().String() + "<===" + string(msg)
		conn.WriteMessage(t, []byte(res))
	}
}

func BinanceWS() {
	url := "wss://stream.binance.com:9443/ws/btcusdt@ticker"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		fmt.Println("Failed to connect websocket : %+v", err)
		return
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			break
		}

		fmt.Println("===>" + string(msg))
	}
}
