package message_managements

import (
  "fmt"
	"github.com/gobuffalo/buffalo"
  "net/http"
	"github.com/gobuffalo/pop"
	"github.com/DappPocket/easy_chain_lennon/models"
  "github.com/gobuffalo/envy"
)

func MessageList(c buffalo.Context) error{
  // Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
  query_hash := c.Param("query_hash")
  c.Set("query_hash", query_hash)

  transactions := &models.Transactions{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())
  if query_hash != "" {
    q = q.Where("hash = ?", query_hash)
  }
  // Retrieve all Transactions from the DB
  if err := q.Order("timestamp DESC").All(transactions); err != nil {
    return err
  }

  c.Set("etherhost", envy.Get("ETHSCAN_PERFIX", "https://etherscan.io"))
  c.Set("transactions", transactions)
  // Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("message_managements/index.html"))
}


func ChangeHiddingMessage(c buffalo.Context) error {
  id := c.Param("id")
  // Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
  record := &models.Transaction{}
  if err := tx.Find(record, id); err != nil {
    c.Flash().Add("error", err.Error())
  }
  record.Hide = !record.Hide
  if err := tx.Update(record); err != nil {
    c.Flash().Add("error", err.Error())
  } else {
    c.Flash().Add("success", "更新成功")
  }

  return c.Redirect(http.StatusSeeOther, c.Request().Referer())
}
