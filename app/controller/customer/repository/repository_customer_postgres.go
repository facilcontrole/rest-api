package repository

import (
	"database/sql"

	"github.com/facilcontrole/rest-api/app/controller/customer"
)

type CustomerRepository struct {
}

func NewPostgresRepository() customer.Repository {
	return &CustomerRepository{}
}

func (c *CustomerRepository) FindAll(conn *sql.DB) (items map[string]interface{}, err error) {
	return
}
