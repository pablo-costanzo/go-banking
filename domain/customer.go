package domain

type (
	Customer struct {
		Id          string
		Name        string
		City        string
		ZipCode     string
		DateofBirth string
		Status      string
	}

	CustomerRepository interface {
		FindAll() ([]Customer, error)
	}
)
