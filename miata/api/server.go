package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"time"

	config "miata/init"
	"miata/model"
)

type Server struct {
	router *gin.Engine
	httpServer *http.Server
	cfg *config.Config
	log *zap.Logger
	store *model.Store


}

func NewServer(lc fx.Lifecycle, cfg *config.Config, log *zap.Logger, store *model.Store) *Server {
	gin.SetMode("debug")
	r := gin.New()
	httpServer := &http.Server{
		Addr:         ":80",
		Handler:      r,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	server := Server{
		router: r,
		httpServer: httpServer,
		cfg: cfg,
		log: log,
		store: store,
	}
	server.registerRoutes()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Stop()
			return nil
		},
	})
	return &server
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	fmt.Println("shutting down server")
	return s.httpServer.Shutdown(context.Background())
}

func (s *Server) registerRoutes() {
	publicGroup := s.router.Group("")
	publicGroup.GET("", s.ProcessRootQuery)
	s.registerEventRoutes(publicGroup)
	s.registerItemRoutes(publicGroup)
}

func (s *Server) ProcessRootQuery(c *gin.Context) {
	c.JSON(
		200, map[string]string{"msg": "ok"},
	)
	return
}