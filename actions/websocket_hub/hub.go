package websocket_hub

import (
  "fmt"
  "math/rand"
  log "github.com/sirupsen/logrus"
)
// Hub...
type Hub struct {
  clients map[*Client]bool
  broadcast chan *Message
  // send2id chan *Message
  register chan *Client
  unregister chan *Client
}

func NewHub() *Hub {
  return &Hub{
    broadcast: make(chan *Message),
    register: make(chan *Client),
    unregister: make(chan *Client),
    clients: make(map[*Client]bool),
  }
}


func (h *Hub) Run(){
  for {
    select {
    // websocket connected
    case client := <-h.register:
      log.Infof("Registered client uid: %s", client.uid)
      h.clients[client] = true
    // websocket disconnected
    case client := <-h.unregister:
      if _, ok := h.clients[client]; ok {
        delete(h.clients, client)
        close(client.buffer)
        log.Infof("Unregistered client uid: %s", client.uid)
      }
    case message := <-h.broadcast:
      for client := range h.clients {
        if !message.IsAvailbeFor(client){
          continue
        }
        data, err := formatMessage(message)
        if err != nil {
          log.Error(err.Error())
          continue
        }
        select {
        case client.buffer <-data:
        default:
          close(client.buffer)
          delete(h.clients, client)
        }
      }
    // case message := <-h.send2id:
    //   for client, status := range h.clients {
    //     if status && client.uid == message.ToUID && message.IsAvailbeFor(client) {
    //       data, err := formatMessage(message)
    //       if err != nil {
    //         log.Error(err.Error())
    //         continue
    //       }
    //       select {
    //       case client.buffer <-data:
    //       default:
    //         close(client.buffer)
    //         delete(h.clients, client)
    //       }
    //     }
    //   }
    }
  }
}

// Authorize ...
func (h *Hub) Authorize(token string) (string, string, error) {
  var uid string
  var name string
  var err error
  uid = fmt.Sprintf("%v", rand.Intn(100000000))
  name = fmt.Sprintf("tclient_%v", uid)
  return uid, name, err
}
