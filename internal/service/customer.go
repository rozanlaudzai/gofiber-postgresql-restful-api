package service

import (
	"context"

	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(cr domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: cr,
	}
}

func (cs *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := cs.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, customer := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   customer.ID,
			Code: customer.Code,
			Name: customer.Name,
		})
	}
	return customerData, nil
}
