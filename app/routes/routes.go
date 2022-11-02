package routes

import (
	v1 "github.com/MdSohelMia/gotipath-kit/app/routes/v1"
	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(v1.NewCustomerRoutes),
	fx.Provide(v1.NewAuthRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	customerRoute v1.CustomerRoutes,
	authRoutes v1.AuthRoutes,
) Routes {
	return Routes{
		customerRoute,
		authRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
