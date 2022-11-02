package v1

import (
	"fmt"

	"github.com/MdSohelMia/gotipath-kit/app/controllers/api/auth"
	"github.com/MdSohelMia/gotipath-kit/core"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger         core.Logger
	handler        core.RequestHandler
	authController auth.AuthController
}

// Setup user routes
func (s AuthRoutes) Setup() {
	s.logger.Info("Setting up routes")
	auth := s.handler.Gin.Group(fmt.Sprintf("%s/auth", RoutePrefix))
	{
		auth.POST("/login", s.authController.Login)
		auth.POST("/register", s.authController.Register)
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler core.RequestHandler,
	authController auth.AuthController,
	logger core.Logger,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
