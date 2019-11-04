package actions
import (
	"net/http"
	"github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/envy"
  log "github.com/sirupsen/logrus"
  "strings"
  "time"
	"github.com/gobuffalo/uuid"
	"fmt"
  "encoding/json"
  sq "github.com/DappPocket/easy_chain_lennon/actions/websocket_hub/sharedq"
  "github.com/ethereum/go-ethereum/common/hexutil"
  md "github.com/DappPocket/easy_chain_lennon/models"
  ah "github.com/DappPocket/easy_chain_lennon/api_helper"
)


type InsertNewTransactionInput struct {
  Hash string `json:"hash"`
	Input string `json:"input"`
	DecoedInput string `json:"decode_input"`
}
func ForceInsertTransaction(c buffalo.Context) error {
	inputs := InsertNewTransactionInput{}
	if err := c.Bind(&inputs); err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{"error": err.Error()}))
	}
	dt, _ := hexutil.Decode(inputs.Input)
	if string(dt) == inputs.DecoedInput {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{"error": ""}))
	}
	data := ah.InfurEthBgResult{
		Hash: inputs.Hash,
		Input: inputs.Input,
	}
	insertOneRecord(data)
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{"msg": "ok"}))
}

func InsertNewTransaction(c buffalo.Context) error {
  inputs := InsertNewTransactionInput{}
  if err := c.Bind(&inputs); err != nil {
    return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{"error": err.Error()}))
  }
  log.Info("hash string", inputs.Hash)
  d := ah.GetOnlyTransaction(inputs.Hash)
  vaildHash := checkToAddr(d)
  if vaildHash {
    data := d.(ah.InfuraOuput)
    // insert new record & notices fronetend update it!
    insertOneRecord(data.Result)
    return c.Render(http.StatusOK, r.JSON(map[string]interface{}{"data": data.Result}))
  } else {
    return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{"error": "not vaild & bad requests"}))
  }
}

func checkToAddr(input interface{}) bool {
  toAddr, err := envy.MustGet("EtherTargetAddr")
  if err != nil {
    log.Error(err.Error())
    return false
  }
  switch input.(type) {
  case string:
    return false
  case ah.InfuraOuput:
    o := input.(ah.InfuraOuput)
    toAddr = strings.ToUpper(toAddr)
    toAddr2 := strings.ToUpper(o.Result.To)
    return toAddr == toAddr2
  default:
    return false
  }
}


func insertOneRecord(o ah.InfurEthBgResult) {
  // bint := hexutil.MustDecodeBig(o.BlockNumber)
  record := md.Transaction{
    BlockNumber:       0,
    Timestamp:         time.Now(),
    Hash:              o.Hash,
    Nonce:             0,
    BlockHash:         o.BlockHash,
    FormAddr:          o.From,
    ToAddr:            o.To,
    Value:             o.Value,
    Gas:               o.Gas,
    GasPrice:          o.GasPrice,
    IsError:           0,
    Input:             o.Input,
    CumulativeGasUsed: "",
    GasUsed:           "",
  }

  newData := md.Transactions{}
	// recordtmp := &md.Transaction{}
	var cid string
	err := md.DB.RawQuery(fmt.Sprintf("select id from transactions where hash = '%s'", record.Hash)).First(&cid)
	if err != nil {
		log.Error(err.Error())
	}
	// err := md.DB.Where("hash = ?", x.Hash).First(&recordtmp)
	if err == nil && cid != "" {
		uid, _ := uuid.FromString(cid)
		record.ID = uid
		if varee, err := md.DB.ValidateAndUpdate(&record); len(varee.Errors) != 0 || err != nil {
			if err != nil {
				log.Error(err)
			}
			if len(varee.Errors) != 0 {
				log.Error("s" + varee.String() + "s")
			}
		}
	} else {
		if err := md.DB.Create(&record); err == nil {
			newData = append(newData, record)
		}
	}
  if len(newData) != 0 {
    datastr, err := json.Marshal(newData)
    if err != nil {
      log.Error(err.Error())
    }
    sq.WebsocketQueue<-sq.MessageData{
      Topic: "datafeed",
      Data: string(datastr),
    }
  }
}
