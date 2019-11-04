package api_helper

import (
  "github.com/gobuffalo/envy"
  log "github.com/sirupsen/logrus"
)

// var APIDOMAIN string = "https://api.etherscan.io"
var APIDOMAIN string = ""
var APIKEY string = envy.Get("EtherScanApiKey", "")

func init() {
  ethhost, err := envy.MustGet("ETHSCAN_API_HOST")
  if err != nil {
    log.Fatalln("etherscan host not set")
  }
  APIDOMAIN = ethhost
  log.SetLevel(log.WarnLevel)
}
