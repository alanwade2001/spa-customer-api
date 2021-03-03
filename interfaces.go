package main

import "github.com/gin-gonic/gin"

// CustomerAPI i
type CustomerAPI interface {
	CreateCustomer(*gin.Context)
	GetCustomer(*gin.Context)
	GetCustomers(*gin.Context)
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
	CreateCustomer(c *Customer) (*Customer, error)
	GetCustomer(id string) (*Customer, error)
	GetCustomers() (*Customers, error)
}

// ConfigAPI si
type ConfigAPI interface {
	Load() error
}
