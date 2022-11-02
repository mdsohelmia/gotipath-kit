package v1

import (
	"github.com/MdSohelMia/gotipath-kit/app/controllers/api/customer"
	"github.com/MdSohelMia/gotipath-kit/core"
)

// CustomerRoutes struct
type CustomerRoutes struct {
	logger             core.Logger
	handler            core.RequestHandler
	customerController customer.CustomerController
}

// Setup Customer routes
func (s CustomerRoutes) Setup() {
	s.logger.Info("Setting up routes")

	v1 := s.handler.Gin.Group("/api/v1")

	v1.GET("/customers", s.customerController.Index)
	v1.POST("/customers", s.customerController.Store)
	v1.GET("/customers/:id", s.customerController.Show)
	v1.PUT("/customers/:id", s.customerController.Update)
	v1.DELETE("/customers/:id", s.customerController.Destroy)
}

// NewCustomerRoutes creates new Customer controller
func NewCustomerRoutes(logger core.Logger, handler core.RequestHandler, customerController customer.CustomerController) CustomerRoutes {
	return CustomerRoutes{
		handler:            handler,
		logger:             logger,
		customerController: customerController,
	}
}
