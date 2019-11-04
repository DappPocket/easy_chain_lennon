package actions

import (
	"strconv"
	"net/http"
	"github.com/gobuffalo/buffalo"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	ah "github.com/DappPocket/easy_chain_lennon/api_helper"
	sq "github.com/DappPocket/easy_chain_lennon/actions/websocket_hub/sharedq"
)

// TriggerQueryQuery default implementation.
func TriggerQueryQuery(c buffalo.Context) error {
	res := ah.GetTransactions()
	ah.CookToTranactionResult(res.Result)
	withResult := c.Param("raw")
	if withResult != "" {
		return c.Render(200, r.JSON(map[string]interface{}{"msg": "ok", "data": res.Result}))
	} else {
		return c.Render(200, r.JSON(map[string]interface{}{"msg": "ok"}))
	}
}

// TriggerQueryQuery default implementation.
func TriggerQueryQueryWithLastNBlock(c buffalo.Context) error {
	lastestBlockNumberStr := ah.GetBlockNumber()
	lastestBnumber, err := strconv.Atoi(lastestBlockNumberStr)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"error": "please retry again."}))
	}
	res := ah.GetTransactions(lastestBnumber-100, lastestBnumber+20)
	ah.CookToTranactionResult(res.Result)
	withResult := c.Param("raw")
	if withResult != "" {
		return c.Render(200, r.JSON(map[string]interface{}{"msg": "ok", "data": res.Result}))
	} else {
		return c.Render(200, r.JSON(map[string]interface{}{"msg": "ok"}))
	}
}

// TriggerQueryQuery default implementation.
func TriggerQueryQueryWithLastNBlockAndWSPushBack(c buffalo.Context) error {
	lastestBlockNumberStr := ah.GetBlockNumber()
	lastestBnumber, err := strconv.Atoi(lastestBlockNumberStr)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"error": "please retry again."}))
	}
	res := ah.GetTransactions(lastestBnumber-100, lastestBnumber+20)
	_, newdata := ah.CookToTranactionResult(res.Result)
	if newdata != nil && len(*newdata) != 0{
		datastr, err := json.Marshal(newdata)
		if err != nil {
			log.Error(err.Error())
		}
		sq.WebsocketQueue<-sq.MessageData{
			Topic: "datafeed",
			Data: string(datastr),
		}
	}
	withResult := c.Param("raw")
	if withResult != "" {
		return c.Render(200, r.JSON(map[string]interface{}{"msg": "ok", "data": newdata}))
	} else {
		return c.Render(200, r.JSON(map[string]interface{}{"msg": "ok"}))
	}
}
