package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Auction struct {
	name string
}

type Bidder struct {
	name string
}

type Ad struct {
	AdId     string
	BidPrice uint64
	IsOpen   bool
}

type AdRequest struct {
	AdPlacementId string
	AdSlot        bool
}

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/hello",
			Handler: func(c echo.Context) error {
				return c.String(200, "We be go-ing...")
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
