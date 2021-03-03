package main

import "github.com/gin-gonic/gin"

// RegisterService s
type RegisterService struct {
	Router      *gin.Engine
	customerAPI CustomerAPI
}

// NewRegisterService f
func NewRegisterService(router *gin.Engine, customerAPI CustomerAPI) RegisterAPI {

	service := RegisterService{router, customerAPI}
	return service

}

// Register f
func (rs RegisterService) Register() error {
	rs.Router.POST("/customers", rs.customerAPI.CreateCustomer)
	rs.Router.GET("/customers", rs.customerAPI.GetCustomers)
	rs.Router.GET("/customers/:id", rs.customerAPI.GetCustomer)

	return nil
}
