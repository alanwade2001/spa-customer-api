package routers

import (
	"net/http"

	"github.com/alanwade2001/spa-customer-api/types"
	"github.com/gin-gonic/gin"
)

// CustomerRouter s
type CustomerRouter struct {
	serviceAPI types.ServiceAPI
}

// NewCustomerRouter f
func NewCustomerRouter(serviceAPI types.ServiceAPI) types.CustomerAPI {

	customerAPI := CustomerRouter{serviceAPI}

	return &customerAPI
}

// CreateCustomer f
func (cr *CustomerRouter) CreateCustomer(ctx *gin.Context) {
	customer := new(types.Customer)

	if err := ctx.BindJSON(customer); err != nil {

		ctx.IndentedJSON(http.StatusUnprocessableEntity, err)

	} else if c1, err := cr.serviceAPI.CreateCustomer(customer); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusCreated, c1)
	}

}

// GetCustomer f
func (cr *CustomerRouter) GetCustomer(ctx *gin.Context) {
	customerID := ctx.Param("id")
	if customer, err := cr.serviceAPI.GetCustomer(customerID); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, customer)
	}
}

// GetCustomers f
func (cr *CustomerRouter) GetCustomers(ctx *gin.Context) {
	if customers, err := cr.serviceAPI.GetCustomers(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, customers)
	}
}
