package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/facilcontrole/app/controller/customer"
	"github.com/facilcontrole/rest-api/app/models"
)

type CustomerRepository struct {
}

func NewPostgresRepository() customer.Repository {
	return &CustomerRepository{}
}

func (c *CustomerRepository) FindAll(conn *sql.DB) (items map[string]interface{}, err error) {
	query, err := conn.Query("SELECT items, id FROM customer")
	items = make(map[string]interface{})
	customers := []models.Customer{}
	customerItem := models.CustomerItems{}
	var jsondata string

	for query.Next() {
		customer := models.Customer{}
		query.Scan(&jsondata, &customer.ID, &customer.Name)
		err = json.Unmarshal([]byte(jsondata), &customerItem)

		if err != nil {
			return nil, err
		}

		customer.Items = customerItem
		customers = append(customers, customer)
	}

	items = map[string]interface{}{}

	items["data"] = customers
	items["total"] = len(customers)
	return
}
