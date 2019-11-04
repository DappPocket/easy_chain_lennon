package shareq

type MessageData struct {
  Topic string `json:"topic"`
  Data string `json:"data"`
}

var WebsocketQueue chan MessageData

func init() {
  WebsocketQueue = make(chan MessageData, 50)
}
