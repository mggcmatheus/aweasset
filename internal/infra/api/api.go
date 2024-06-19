package api

import (
	"aweasset/internal/infra/api/router"
	"aweasset/internal/infra/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func InitServer(cfg *config.Config) {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	RegisterRoutes(server, cfg)

	// Listar endpoints
	routes := server.Routes()
	fmt.Println("Endpoints disponíveis:")
	for _, route := range routes {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}

	err := server.Start(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal("Porta em uso: ", cfg.Server.Port)
	}
}

func RegisterRoutes(s *echo.Echo, cfg *config.Config) {

	v1 := s.Group("/api/v1")
	{
		testRouter := v1.Group("/test")
		// Test
		router.TestRouter(testRouter)
	}
}
