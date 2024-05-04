package service

import (
	"github.com/pablo-costanzo/go-banking/domain"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, error) {
	return s.repo.FindAll()
}
