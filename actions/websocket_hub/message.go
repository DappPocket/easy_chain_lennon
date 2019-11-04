package websocket_hub

import (
  "time"
  "encoding/json"
)

type Message struct {
  Source string `json:"soruce"`
  FormUID string `json:"form_uid,omitempty"`
  FormName string `json:"form_name,omitempty"`
  ToUID string `json:"to_uid,omitempty"`
  ReceivedAt time.Time `json:"received_at"`
  Topic string `json:"topic"`
  Data interface{} `json:"data"`
}

func formatMessage(message *Message) ([]byte, error){
  return json.Marshal(message)
}

func (message *Message) IsAvailbeFor(client *Client) bool {
  // sender should eq to message sender
  return message.FormUID != "" && message.FormUID == client.uid ||
  // touid should not empty & toid not eq to sender
    message.ToUID != "" && message.ToUID != client.uid
}
