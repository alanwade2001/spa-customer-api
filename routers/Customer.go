package routers

import (
	"io"
	"net/http"

	"github.com/alanwade2001/spa-customer-api/models/generated"
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

	customer := new(generated.CustomerModel)

	if err := func() error {
		if bytes, err := io.ReadAll(ctx.Request.Body); err != nil {

			return err

		} else if err := customer.UnmarshalJSON(bytes); err != nil {

			return err

		}

		return nil

	}(); err != nil {

		ctx.String(http.StatusUnprocessableEntity, err.Error())

	} else if _, err := cr.serviceAPI.CreateCustomer(customer); err != nil {

		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.IndentedJSON(http.StatusCreated, customer)

}

// GetCustomer f
func (cr *CustomerRouter) GetCustomer(ctx *gin.Context) {
	customerID := ctx.Param("id")
	if customer, err := cr.serviceAPI.GetCustomer(customerID); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else if customer == nil {
		ctx.Status(http.StatusNotFound)
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
