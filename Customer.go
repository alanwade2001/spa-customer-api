package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerRouter s
type CustomerRouter struct {
	repositoryAPI RepositoryAPI
}


// NewCustomerRouter f
func NewCustomerRouter(repositoryAPI RepositoryAPI) CustomerAPI {

	customerAPI := CustomerRouter{repositoryAPI}

	return &customerAPI
}

// CreateCustomer f
func (cr *CustomerRouter) CreateCustomer(ctx *gin.Context) {
	customer := new(Customer)

	if err := ctx.BindJSON(customer); err != nil {

		ctx.IndentedJSON(http.StatusUnprocessableEntity, err)

	} else if c1, err := cr.repositoryAPI.CreateCustomer(customer); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusCreated, c1)
	}

}

// GetCustomer f
func (cr *CustomerRouter) GetCustomer(ctx *gin.Context) {
	customerID := ctx.Param("id")
	if customer, err := cr.repositoryAPI.GetCustomer(customerID); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, customer)
	}
}

// GetCustomers f
func (cr *CustomerRouter) GetCustomers(ctx *gin.Context) {
	if customers, err := cr.repositoryAPI.GetCustomers(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, customers)
	}
}

