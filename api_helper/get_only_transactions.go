package api_helper

import (
  gq "github.com/parnurzeal/gorequest"
  "github.com/gobuffalo/envy"
  "fmt"
  "encoding/json"
)

type InfurEthBgResult struct {
  BlockHash string `json:"blockHash"`
  BlockNumber string `json:"blockNumber"`
  From string `json:"from"`
  To  string `json:"to"`
  Gas string `json:"gas"`
  GasPrice  string `json:"gasPrice"`
  Hash string `json:"hash"`
  Input string `json:"input"`
  Nonce string `json:"nonce"`
  Value string `json:"value"`
  TransactionIndex string `json:"transactionIndex"`
  R string `json:"r"`
  S string `json:"s"`
  V string `json:"v"`
}

type InfuraOuput struct {
  JSONRPC string `json:"jsonrpc"`
  ID int `json:"id"`
  Result InfurEthBgResult `json:"result"`
}

func GetOnlyTransaction(transactionHash string) interface{}{
  inf, err := envy.MustGet("INFRA_API_HOST")
  if err != nil {
    return err.Error()
  }
  req := gq.New()
  req = req.Post(inf)
  req.AppendHeader("Content-Type", "application/json")
  resp, body , errs := req.SendStruct(map[string]interface{}{
    "jsonrpc":"2.0",
    "method":"eth_getTransactionByHash",
    "params": []string{transactionHash},
    "id": 1,
  }).End()
  if resp.StatusCode != 0 && errs != nil {
    return fmt.Sprintf("%v", errs)
  }
  output := InfuraOuput{}
  json.Unmarshal([]byte(body), &output)
  return output
}
