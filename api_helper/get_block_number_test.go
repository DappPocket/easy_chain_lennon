package api_helper

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestGetBlock(t *testing.T) {
  Convey("test get block number", t, func(){
    result := GetBlockNumber()
    So(result, ShouldNotBeEmpty)
  })
}
