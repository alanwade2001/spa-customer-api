package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerRouter s
type CustomerRouter struct {
	repositoryAPI RepositoryAPI
}

// RegisterService s
type RegisterService struct {
	Router      *gin.Engine
	customerAPI CustomerAPI
}

// NewRegisterAPI f
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

// NewServer f
func NewServer(router *gin.Engine, registerAPI RegisterAPI, configAPI ConfigAPI) ServerAPI {

	return &Server{router, registerAPI, configAPI}
}

// Server s
type Server struct {
	Router      *gin.Engine
	registerAPI RegisterAPI
	configAPI   ConfigAPI
}

// Run f
func (s *Server) Run() error {
	if err := s.configAPI.Load(); err != nil {
		return err
	}

	if err := s.registerAPI.Register(); err != nil {
		return err
	}

	return s.Router.Run()
}
