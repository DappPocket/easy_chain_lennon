package websocket_hub

import (
  "net/http"
  "time"
  "fmt"
  "github.com/pkg/errors"
  log "github.com/sirupsen/logrus"
  "github.com/gobuffalo/buffalo"
  "github.com/gorilla/websocket"
  "github.com/gobuffalo/buffalo/render"
  sq "github.com/DappPocket/easy_chain_lennon/actions/websocket_hub/sharedq"
)

const (
  maxMessageSize = 512
  writeWait      = 10 * time.Second
  pongWait       = 60 * time.Second
  pingPeriod     = (pongWait * 9) / 10
)

var defaultUpgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}

var client *Client
func NoticeNewSocketHandler(hub *Hub, c buffalo.Context) error {
  log.Infof("Handling Websocket connection from %v", c.Request().RemoteAddr)
  uid, name, err := hub.Authorize(c.Request().URL.Query().Get("token"))
  if err != nil {
    c.Response().WriteHeader(403)
    log.Error(err.Error())
    errors.WithStack(err)
  }
  ws, err := websocket.Upgrade(c.Response(), c.Request(), c.Response().Header(), 1024, 1024)
  if err != nil {
    if _, ok := err.(websocket.HandshakeError); !ok {
      errors.WithStack(err)
    }
  }
  log.Debug(uid, name, ws)
  client = &Client{hub: hub, conn: ws, uid: uid, name: name, buffer: make(chan []byte, 256)}
  client.hub.register <- client
  go client.Reader()
  client.Writer()
  client.hub.unregister <- client
  client.conn.Close()
  return nil
}

var r *render.Engine
// TestPush
func TestPush(c buffalo.Context) error {
	sq.WebsocketQueue<-sq.MessageData{
    Topic: "test",
    Data: `[{"msg": "hello"}]`,
  }
  return c.Render(200, r.JSON(map[string]string{"msg": "1"}))
}

func WSClients(c buffalo.Context) error {
  var clientList []string
  if(client != nil){
    for k, v := range client.hub.clients{
      clientList = append(clientList, fmt.Sprintf("id: %s, name: %s, enable: %v", k.uid, k.name, v))
    }
  }
  return c.Render(200, r.JSON(map[string]interface{}{"clients": clientList}))
}
