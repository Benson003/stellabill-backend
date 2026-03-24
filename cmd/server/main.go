package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"stellarbill-backend/internal/config"
	"stellarbill-backend/internal/middleware"
	"stellarbill-backend/internal/routes"
)

func main() {
	cfg := config.Load()
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := newRouter()

	addr := ":" + cfg.Port
	log.Printf("Stellarbill backend listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(
		middleware.Recovery(log.Default()),
		middleware.RequestID(),
		middleware.Logging(log.Default()),
		middleware.CORS("*"),
		middleware.RateLimit(middleware.NewRateLimiter(60, time.Minute)),
	)
	routes.Register(router)
	return router
}
