package customer

import (
	"net/http"

	"github.com/MdSohelMia/gotipath-kit/core"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	logger core.Logger
	db     core.Database
}

func NewCustomerController(logger core.Logger, db core.Database) CustomerController {
	return CustomerController{
		logger: logger,
		db:     db,
	}
}
func (customer CustomerController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "index"})
}

func (customer CustomerController) Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Store"})
}

func (customer CustomerController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Update"})
}

func (customer CustomerController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Show"})
}

func (customer CustomerController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Destroy"})
}
