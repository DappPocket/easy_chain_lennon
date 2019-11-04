package accounts

import (
  "github.com/gobuffalo/envy"
  "github.com/gobuffalo/buffalo"
  "net/http"
)


func SessionCheck(next buffalo.Handler) buffalo.Handler {
  loginsession, err := envy.MustGet("SESSION")
  return func(c buffalo.Context) error {
    if err != nil {
      return err
    }
    currentauth := c.Session().Get("authentication")
    if currentauth == nil || currentauth.(string) != loginsession {
      return c.Redirect(http.StatusSeeOther, "/loginpage")
    }
    // do some work before calling the next handler
    err = next(c)
    // do some work after calling the next handler
    return err
  }
}
