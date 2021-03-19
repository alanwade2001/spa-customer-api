package types

import (
	models "github.com/alanwade2001/spa-customer-api/models/generated"
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
	CreateCustomer(c *models.CustomerModel) (*models.CustomerModel, error)
	GetCustomer(id string) (*models.CustomerModel, error)
	GetCustomers() (*[]models.CustomerModel, error)

	FindCustomerByEmail(id string) (*models.CustomerModel, error)
}

// ConfigAPI si
type ConfigAPI interface {
	Load(configPath string) error
}

type ServiceAPI interface {
	CreateCustomer(c *models.CustomerModel) (*models.CustomerModel, error)
	GetCustomer(id string) (*models.CustomerModel, error)
	GetCustomers() (*[]models.CustomerModel, error)

	FindCustomerByEmail(id string) (*models.CustomerModel, error)
}
