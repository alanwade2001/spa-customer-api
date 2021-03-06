package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerSearchRouter s
type CustomerSearchRouter struct {
	repositoryAPI RepositoryAPI
}

// NewCustomerSearchRouter f
func NewCustomerSearchRouter(repositoryAPI RepositoryAPI) CustomerSearchAPI {
	return CustomerSearchRouter{
		repositoryAPI: repositoryAPI,
	}
}

// FindCustomerByEmail f
func (cs CustomerSearchRouter) FindCustomerByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if customer, err := cs.repositoryAPI.FindCustomerByEmail(email); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, customer)
	}
}
