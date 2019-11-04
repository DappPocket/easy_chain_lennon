package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"
	"github.com/rs/cors"
	"net/http"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
	"github.com/DappPocket/easy_chain_lennon/models"
	wh "github.com/DappPocket/easy_chain_lennon/actions/websocket_hub"
	acct "github.com/DappPocket/easy_chain_lennon/actions/accounts"
	magc "github.com/DappPocket/easy_chain_lennon/actions/message_managements"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
      Env:         ENV,
			SessionName: "_easychainlife_bg_session",
      PreWares: []buffalo.PreWare{cors.New(cors.Options{
          AllowedOrigins: []string{"*"},
          AllowedMethods: []string{
            http.MethodHead,
            http.MethodGet,
            http.MethodPost,
            http.MethodPut,
            http.MethodPatch,
            http.MethodDelete,
            http.MethodOptions,
          },
          AllowedHeaders:   []string{"*"},
          AllowCredentials: true,
      }).Handler},
    })

		// Automatically redirect to SSL
    // disable froce ssl
		forcessl := envy.Get("FROCE_SSL", "false")
		if(forcessl != "false"){
			app.Use(forceSSL())
		}

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)
		// enable websocket
		hub := wh.NewHub()
		go hub.Run()
		app.GET("/chain_watch", func(c buffalo.Context) error {
			return wh.NoticeNewSocketHandler(hub, c)
		})
		app.GET("/testpush", wh.TestPush)
		app.Resource("/watch_addresses", WatchAddressesResource{})
		app.Resource("/transactions", TransactionsResource{})
		app.GET("/trigger_query/query", TriggerQueryQuery)
		app.GET("/update_recent_transcations", TriggerQueryQueryWithLastNBlock)
		app.GET("/ws/update_recent_transcations", TriggerQueryQueryWithLastNBlockAndWSPushBack)
		app.GET("/clients", wh.WSClients)
		app.POST("/add_submmited_tx", InsertNewTransaction)
		app.POST("/api/v1/forces_insert_one", ForceInsertTransaction)

		// LoginPage
		app.GET("/loginpage", acct.LoginPage)
		app.POST("/user_login", acct.LoginAction)
		app.GET("/logout", acct.LogOut)
		mag := app.Group("/message_managements")
		mag.Use(acct.SessionCheck)
		mag.GET("/list", magc.MessageList)
		mag.GET("/change_hidding/{id}", magc.ChangeHiddingMessage)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
