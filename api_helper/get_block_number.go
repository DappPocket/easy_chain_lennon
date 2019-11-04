package api_helper

import (
  "fmt"
  "encoding/json"
  "strings"
  "strconv"
  gq "github.com/parnurzeal/gorequest"
)

type GetBlockNumberResp struct {
  Jsonrpc string  `json:"jsonrpc"`
  Id int `json:"id"`
  Result string `json:"result"`
}

// get latest block numbers
func GetBlockNumber() string {
  req := gq.New()
  payload := "module=proxy&action=eth_blockNumber"
  _, body, _ := req.Get(fmt.Sprintf("%s/api?%s&apikey=%s", APIDOMAIN, payload, APIKEY)).End()
  outputs := GetBlockNumberResp{}
  err := json.Unmarshal([]byte(body), &outputs)
  if err != nil {
    fmt.Println(err.Error())
    return ""
  }
  hexblock := strings.Replace(outputs.Result, "0x", "", 1)
  n, err := strconv.ParseInt(hexblock, 16, 32)
  if err != nil {
    fmt.Println(err.Error())
    return ""
  }
  return fmt.Sprintf("%v", n)
}
