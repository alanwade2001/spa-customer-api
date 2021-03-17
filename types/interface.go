package types

import (
	"github.com/alanwade2001/spa-customer-api/models/generated"
	"github.com/gin-gonic/gin"
)

// CustomerAPI i
type CustomerAPI interface {
	CreateCustomer(*gin.Context)
	GetCustomer(*gin.Context)
	GetCustomers(*gin.Context)
}

// CustomerSearchAPI i
type CustomerSearchAPI interface {
	FindCustomerByEmail(*gin.Context)
}

// ServerAPI i
type ServerAPI interface {
	Run() error
}

// RegisterAPI i
type RegisterAPI interface {
	Register() error
}

// RepositoryAPI i
type RepositoryAPI interface {
	CreateCustomer(c *generated.CustomerModel) (*generated.CustomerModel, error)
	GetCustomer(id string) (*generated.CustomerModel, error)
	GetCustomers() (*[]generated.CustomerModel, error)

	FindCustomerByEmail(id string) (*generated.CustomerModel, error)
}

// ConfigAPI si
type ConfigAPI interface {
	Load() error
}

type ServiceAPI interface {
	CreateCustomer(c *Customer) (*Customer, error)
	GetCustomer(id string) (*Customer, error)
	GetCustomers() (*Customers, error)

	FindCustomerByEmail(id string) (*Customer, error)
}
