package services

import (
	models "github.com/alanwade2001/spa-customer-api/models/generated"
	"github.com/alanwade2001/spa-customer-api/types"
)

// Service s
type Service struct {
	repository types.RepositoryAPI
}

// NewService func
func NewService(repository types.RepositoryAPI) types.ServiceAPI {

	return &Service{
		repository: repository,
	}
}

func (Service *Service) CreateCustomer(c *models.CustomerModel) (*models.CustomerModel, error) {
	if _, err := Service.repository.CreateCustomer(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (Service *Service) GetCustomer(id string) (customer *models.CustomerModel, err error) {

	if customer, err = Service.repository.GetCustomer(id); err != nil {
		return nil, err
	}

	return customer, nil
}
func (Service *Service) GetCustomers() (customers *[]models.CustomerModel, err error) {
	if customers, err = Service.repository.GetCustomers(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (Service *Service) FindCustomerByEmail(email string) (customer *models.CustomerModel, err error) {
	if customer, err = Service.repository.FindCustomerByEmail(email); err != nil {
		return nil, err
	}

	return customer, nil
}
