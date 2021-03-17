package routers

import (
	"github.com/alanwade2001/spa-customer-api/types"
	"github.com/gin-gonic/gin"
)

// RegisterService s
type RegisterService struct {
	Router            *gin.Engine
	customerAPI       types.CustomerAPI
	customerSearchAPI types.CustomerSearchAPI
}

// NewRegisterService f
func NewRegisterService(router *gin.Engine, customerAPI types.CustomerAPI, customerSearchAPI types.CustomerSearchAPI) types.RegisterAPI {

	service := RegisterService{router, customerAPI, customerSearchAPI}
	return service

}

// Register f
func (rs RegisterService) Register() error {
	rs.Router.POST("/customers", rs.customerAPI.CreateCustomer)
	rs.Router.GET("/customers", rs.customerAPI.GetCustomers)
	rs.Router.GET("/customers/:id", rs.customerAPI.GetCustomer)

	rs.Router.GET("/customer/search/email/:email", rs.customerSearchAPI.FindCustomerByEmail)

	return nil
}
