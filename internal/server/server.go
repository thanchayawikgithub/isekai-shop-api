package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	adminRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/repositories"
	oauth2Controllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/controllers"
	oauth2Services "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/services"
	playerRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/routes"
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	app := echo.New()

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
	mw := middlewares.NewMiddleware(s.app, s.conf.Server)
	mw.RegisterMiddleWares()

	s.app.GET("/v1/health", s.healthCheck)
	router := routes.NewRouter(s.app, s.db, s.app.Logger, s.conf)

	playerRepo := playerRepositories.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepo := adminRepositories.NewAdminRepositoryImpl(s.db, s.app.Logger)
	oauth2Service := oauth2Services.NewGoogleOAuth2Service(playerRepo, adminRepo)
	oauth2Controller := oauth2Controllers.NewGoogleOAuth2Controller(oauth2Service, s.conf.OAuth2, s.app.Logger)
	authMiddleWare := middlewares.NewAuthMiddleware(oauth2Controller, s.conf.OAuth2, s.app.Logger)

	router.RegisterRoutes(authMiddleWare)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

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

func (s *echoServer) gracefullyShutdown(quitCh chan os.Signal) {
	ctx := context.Background()

	<-quitCh
	s.app.Logger.Info("Shutting down server")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}
