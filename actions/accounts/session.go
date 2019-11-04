package accounts

import (
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/envy"
)

func LoginPage(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("accounts/loginpage.html"))
}

type LoginActionInput struct {
  Username string `json:"username" form:"username"`
  Password string `json:"password" form:"password"`
}
func LoginAction(c buffalo.Context) error {
  inputs := LoginActionInput{}
  if err := c.Bind(&inputs); err != nil {
    return err
  }
  sname, err := envy.MustGet("USERNAME")
  if err != nil {
    return err
  }
  spasswd, err := envy.MustGet("PASSWORD")
  if err != nil {
    return err
  }
  if sname == inputs.Username && spasswd == spasswd {
    loginsession, err := envy.MustGet("SESSION")
    if err != nil {
      return err
    }
    c.Session().Set("authentication", loginsession)
    c.Flash().Add("success", "登入成功")
    return c.Redirect(http.StatusMovedPermanently, "/message_managements/list")
    // return c.Render(http.StatusOK, r.JSON(map[string]interface{}{"status": "pass", "session": loginsession, "input": c.Params(), "bind": inputs}))
  } else {
    c.Flash().Add("error", "帳號or密碼不正確")
    return c.Redirect(http.StatusSeeOther, "/loginpage")
    // return c.Render(http.StatusOK, r.JSON(map[string]interface{}{"error": "nomatch", "input": c.Params(), "bind": inputs}))
  }
}

func LogOut(c buffalo.Context) error {
  c.Session().Delete("authentication")
  return c.Redirect(http.StatusSeeOther, "/loginpage")
}
