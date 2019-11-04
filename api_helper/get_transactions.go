package api_helper

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/uuid"
	md "github.com/DappPocket/easy_chain_lennon/models"
	gq "github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type GetTransactionsResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  []GTRResult `json:"result"`
}

type GTRResult struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string
	Hash              string `json:"hash"`
	Nonce             string
	blockHash         string
	From              string
	To                string
	Value             string
	Gas               string
	GasPrice          string
	IsError           string
	Input             string
	CumulativeGasUsed string
	GasUsed           string
}

func GetTransactions(blocknumberset ...int) *GetTransactionsResp {
	startBlock := 0
	endBlock := 99999999
	if len(blocknumberset) == 2 {
		startBlock = blocknumberset[0]
		endBlock = blocknumberset[1]
	}
	req := gq.New()
	targetAddr := "0x99992213Adf6471e52ED09EF47B36Faf7b769600"
	payload := fmt.Sprintf("module=account&action=txlist&address=%s&startblock=%d&endblock=%d&sort=asc", targetAddr, startBlock, endBlock)
	_, body, _ := req.Get(fmt.Sprintf("%s/api?%s&apikey=%s", APIDOMAIN, payload, APIKEY)).End()
	log.SetLevel(log.InfoLevel)
	log.Info("request url", req.Url)
	outputs := GetTransactionsResp{}
	if err := json.Unmarshal([]byte(body), &outputs); err != nil {
		return nil
	}
	return &outputs
}

func ConverGTRToModelTransactions(inputs []GTRResult) []*md.Transaction {
	newData := []*md.Transaction{}
	for _, x := range inputs {
		bk, _ := strconv.Atoi(x.BlockNumber)
		tp, _ := strconv.Atoi(x.TimeStamp)
		nc, _ := strconv.Atoi(x.Nonce)
		ir, _ := strconv.Atoi(x.IsError)
		record := &md.Transaction{
			BlockNumber:       bk,
			Timestamp:         time.Unix(int64(tp), 0),
			Hash:              x.Hash,
			Nonce:             nc,
			BlockHash:         x.blockHash,
			FormAddr:          x.From,
			ToAddr:            x.To,
			Value:             x.Value,
			Gas:               x.Gas,
			GasPrice:          x.GasPrice,
			IsError:           ir,
			Input:             x.Input,
			CumulativeGasUsed: x.CumulativeGasUsed,
			GasUsed:           x.GasUsed,
		}
		record.Message = record.InputString()
		newData = append(newData, record)
	}
	return newData
}

func CookToTranactionResult(inputs []GTRResult) (error, *md.Transactions) {
	newData := md.Transactions{}
	for _, x := range inputs {
		bk, _ := strconv.Atoi(x.BlockNumber)
		tp, _ := strconv.Atoi(x.TimeStamp)
		nc, _ := strconv.Atoi(x.Nonce)
		ir, _ := strconv.Atoi(x.IsError)
		record := md.Transaction{
			BlockNumber:       bk,
			Timestamp:         time.Unix(int64(tp), 0),
			Hash:              x.Hash,
			Nonce:             nc,
			BlockHash:         x.blockHash,
			FormAddr:          x.From,
			ToAddr:            x.To,
			Value:             x.Value,
			Gas:               x.Gas,
			GasPrice:          x.GasPrice,
			IsError:           ir,
			Input:             x.Input,
			CumulativeGasUsed: x.CumulativeGasUsed,
			GasUsed:           x.GasUsed,
		}
		record.Message = record.InputString()
		// recordtmp := &md.Transaction{}
		var cid string
		err := md.DB.RawQuery(fmt.Sprintf("select id from transactions where hash = '%s' and hide != %s", x.Hash, "true")).First(&cid)
		if err != nil {
			log.Error(err.Error())
		}
		// err := md.DB.Where("hash = ?", x.Hash).First(&recordtmp)
		if err == nil && cid != "" {
			var blockNumber int
			err := md.DB.RawQuery(fmt.Sprintf("select block_number from transactions where hash = '%s'", x.Hash)).First(&blockNumber)
			if err != nil {
				log.Error(err.Error())
				continue
			}
			uid, _ := uuid.FromString(cid)
			record.ID = uid
			if varee, err := md.DB.ValidateAndUpdate(&record); len(varee.Errors) != 0 || err != nil {
				if err != nil {
					log.Error(err)
				}
				if len(varee.Errors) != 0 {
					log.Error(varee.String())
				}
			}
			// when pervious block is empty. push to frontend
			if blockNumber == 0 {
				newData = append(newData, record)
			}
		} else {
			if err := md.DB.Create(&record); err == nil {
				newData = append(newData, record)
			}
		}
	}
	return nil, &newData
}
