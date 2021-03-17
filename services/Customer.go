package services

import (
	"github.com/alanwade2001/spa-customer-api/types"
)

// Service s
type Service struct {
}

// NewService func
func NewService() types.ServiceAPI {

	return &Service{}
}

func (Service *Service) CreateCustomer(c *types.Customer) (*types.Customer, error) {
	return nil, nil
}

func (Service *Service) GetCustomer(id string) (*types.Customer, error) {
	return nil, nil
}
func (Service *Service) GetCustomers() (*types.Customers, error) {
	return nil, nil
}

func (Service *Service) FindCustomerByEmail(id string) (*types.Customer, error) {
	return nil, nil
}
