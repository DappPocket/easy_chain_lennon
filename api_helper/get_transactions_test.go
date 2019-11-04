package api_helper

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  log "github.com/sirupsen/logrus"
)

func TestGetTx(t *testing.T) {
  log.SetLevel(log.WarnLevel)
  Convey("test get transcation list", t, func (){
    result := GetTransactions()
    output := GetTransactionsResp{}
    log.Errorf("%v", output.Result)
    So(result, ShouldNotBeEmpty)
  })
}
