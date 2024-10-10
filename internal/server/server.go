package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"gorm.io/gorm"
)

type echoServer struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app,
			db,
			conf,
		}
	})

	return server
}

func (s *echoServer) Start() {
	s.app.GET("/v1/health", s.healthCheck)
	s.httpListenting()
}

func (s *echoServer) httpListenting() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatal("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
