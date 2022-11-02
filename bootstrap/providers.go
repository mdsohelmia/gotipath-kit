package bootstrap

import (
	"github.com/MdSohelMia/gotipath-kit/app/controllers"
	"github.com/MdSohelMia/gotipath-kit/app/routes"
	"github.com/MdSohelMia/gotipath-kit/core"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	routes.Module,
	core.Module,
	controllers.Module,
)
