package websocket_hub
import (
  "bytes"
  "time"
  "encoding/json"
  log "github.com/sirupsen/logrus"
  "github.com/gorilla/websocket"
  "github.com/pkg/errors"
  sq "github.com/DappPocket/easy_chain_lennon/actions/websocket_hub/sharedq"
)

type Client struct {
  hub * Hub
  conn *websocket.Conn
  uid string
  name string
  buffer chan []byte
}

//Writer...
func(c *Client) Writer(){
  pingTicker := time.NewTicker(pingPeriod)

	defer func() {
		pingTicker.Stop()
	}()

	for {
		select {
		case message, ok := <-c.buffer:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte("["))
			w.Write(message)
			for i := 0; i < len(c.buffer); i++ {
				w.Write([]byte(","))
				w.Write(<-c.buffer)
			}
			w.Write([]byte("]"))
			if err := w.Close(); err != nil {
				return
			}
    case messageData:=<-sq.WebsocketQueue:
      msdata := Message{
        Source: "backend_push",
        FormUID: c.uid,
        FormName: c.name,
        ToUID: "all",
        ReceivedAt: time.Now(),
        Topic: messageData.Topic,
        Data: messageData.Data,
      }
      c.hub.broadcast<-&msdata
      // b, _ := json.Marshal(msdata)
      // c.buffer<-b
		case <-pingTicker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
	 }
  }
}


// Reader ...
func (c *Client) Reader() {
	defer func() {
		// c.hub.unregister <- c
		// c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			errors.WithStack(err)
      log.Errorf("Internal Error with %v", err.Error())
			c.WriteError("Internal Error")
			break
		}
		message := string(bytes.TrimSpace(bytes.Replace(data, []byte{'\n'}, []byte{' '}, -1)))
		if len(message) < 1 {
			c.WriteError("Invalid data")
			break
		}
		messageData := new(sq.MessageData)
		if err := json.Unmarshal([]byte(message), messageData); err != nil {
			log.Errorf("recevied error data: %v", message)
			c.WriteError("Invalid data")
			// break
		}else{
		  log.Infof("recevied data: %v", message)
		}
	}
}

// WriteError ...
func (c *Client) WriteError(message string) {
	messageData := make(map[string]interface{})
	messageData["error"] = message
	// c.hub.send2id <- &Message{
  //   FormUID: "backend",
  //   FormName: "backend",
	// 	Source: "server",
	// 	ToUID:  c.uid,
	// 	Data:   messageData,
	// }
	log.Errorf("WriteError: %v", message)

	c.hub.unregister <- c
}
