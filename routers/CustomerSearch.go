package routers

import (
	"net/http"

	"github.com/alanwade2001/spa-customer-api/types"
	"github.com/gin-gonic/gin"
)

// CustomerSearchRouter s
type CustomerSearchRouter struct {
	repositoryAPI types.RepositoryAPI
}

// NewCustomerSearchRouter f
func NewCustomerSearchRouter(repositoryAPI types.RepositoryAPI) types.CustomerSearchAPI {
	return CustomerSearchRouter{
		repositoryAPI: repositoryAPI,
	}
}

// FindCustomerByEmail f
func (cs CustomerSearchRouter) FindCustomerByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if customer, err := cs.repositoryAPI.FindCustomerByEmail(email); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else if customer == nil {
		ctx.Status(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, customer)
	}
}
