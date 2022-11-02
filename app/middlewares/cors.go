package middlewares

import (
	"github.com/MdSohelMia/gotipath-kit/core"
	"github.com/gin-contrib/cors"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler core.RequestHandler
	logger  core.Logger
	env     core.Env
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(handler core.RequestHandler, logger core.Logger, env core.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// AllowOrigins:     []string{"https://foo.com"},
// AllowMethods:     []string{"PUT", "PATCH"},
// AllowHeaders:     []string{"Origin"},
// ExposeHeaders:    []string{"Content-Length"},
// AllowCredentials: true,
// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	// debug := m.env.Environment == "development"
	m.handler.Gin.Use(cors.New(cors.Config{}))
}
