package controllers

import (
	"github.com/MdSohelMia/gotipath-kit/app/controllers/api/auth"
	"github.com/MdSohelMia/gotipath-kit/app/controllers/api/customer"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(auth.NewAuthController),
	fx.Provide(customer.NewCustomerController),
)
