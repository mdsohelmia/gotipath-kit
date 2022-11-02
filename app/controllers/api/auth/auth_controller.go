package auth

import (
	"github.com/MdSohelMia/gotipath-kit/core"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger core.Logger
	db     core.Database
}

func NewAuthController(logger core.Logger, db core.Database) AuthController {
	return AuthController{
		logger: logger,
		db:     db,
	}
}
func (auth AuthController) Login(c *gin.Context) {
	c.JSON(200, "ok")
}

func (auth AuthController) Register(c *gin.Context) {
	c.JSON(200, "ok")
}
