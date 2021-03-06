package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	gommonlog "github.com/labstack/gommon/log"
	"github.com/ytake/draft/action"
	"github.com/ytake/draft/config"
	"github.com/ytake/draft/record"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	echo.NotFoundHandler = action.HTTPVndErrorResponse
	e := echo.New()
	e.Logger.SetLevel(gommonlog.INFO)
	e.Use(echoMiddleware.RemoveTrailingSlash())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: config.AllowedOrigins(),
	}))
	c := config.NewConfig()
	r, err := config.Parse(c.File)
	if err != nil {
		e.Logger.Fatal()
	}

	h := &action.Handle{
		Documenter: record.NewRecorder(r),
	}
	e.GET("/ping", h.Ping)
	e.POST("/documents", h.AddDocument)
	e.GET("/documents", h.FindDocument)

	// graceful
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	if err := e.Start(fmt.Sprintf(":%s", c.Port)); err != nil {
		e.Logger.Info("shutting down the server")
	}
}
