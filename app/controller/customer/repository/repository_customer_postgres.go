package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/facilcontrole/rest-api/app/controller/customer"
	"github.com/facilcontrole/rest-api/app/models"
)

type CustomerRepository struct {
}

func NewPostgresRepository() customer.Repository {
	return &CustomerRepository{}
}

func (c *CustomerRepository) FindAll(conn *sql.DB) (item map[string]interface{}, err error) {

	query := "SELECT id, items FROM customer"

	rows, err := conn.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	var items []map[string]interface{}

	for rows.Next() {

		var id, _items string

		err = rows.Scan(&id, &_items)

		if err != nil {
			return
		}

		var customer models.CustomerItems

		err = json.Unmarshal([]byte(_items), &customer)

		if err != nil {
			return
		}

		it := map[string]interface{}{
			"id":    id,
			"phone": customer.Phone,
		}

		items = append(items, it)
	}

	item = map[string]interface{}{
		"items": items,
		"total": 1,
	}

	return
}
