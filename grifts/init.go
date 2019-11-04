package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/DappPocket/easy_chain_lennon/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
