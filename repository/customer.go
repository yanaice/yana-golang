package repository

// PORT
type Customer struct {
	CustomerID string `db:"customer_id"`
	Name       string `db:"name"`
	Age        int    `db:"age"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(string) (*Customer, error)
}
