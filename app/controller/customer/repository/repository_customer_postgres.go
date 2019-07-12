package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/facilcontrole/rest-api/app/controller/customer"
	"github.com/facilcontrole/rest-api/app/models"
	uuid "github.com/nu7hatch/gouuid"
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

func (c *CustomerRepository) Storage(name string, body []byte, conn *sql.DB) (id string, err error) {

	query := "INSERT INTO customer(id,name,updated_at,items) VALUES($1,$2,$3,$4)"
	smt, err := conn.Prepare(query)

	if err != nil {
		return
	}

	_id, _ := uuid.NewV4()

	date := time.Now().Format("2006-01-02T15:04:05Z07:00")

	_, err = smt.Exec(_id.String(), name, date, string(body))

	if err != nil {
		return
	}

	id = _id.String()

	return
}

func (c *CustomerRepository) FindById(id string, conn *sql.DB) (items, name string, err error) {
	query := "SELECT items, name FROM customer WHERE id=$1"

	row := conn.QueryRow(query, id)

	var _name sql.NullString

	err = row.Scan(&items, &_name)

	name = _name.String

	return
}

func (c *CustomerRepository) Update(item models.Customer, conn *sql.DB) (err error) {

	query := fmt.Sprintf(`
	UPDATE customer 
			SET 
			  items = jsonb_set(items,'{phone}','"%s"'),
			  name=$1,
			  updated_at=$2
			WHERE id=$3
`, item.Items.Phone)

	smt, err := conn.Prepare(query)

	if err != nil {
		return
	}

	date := time.Now().Format("2006-01-02T15:04:05Z07:00")
	_, err = smt.Exec(item.Name, date, item.ID)

	return
}

func (c *CustomerRepository) Delete(id string, conn *sql.DB) (err error) {

	query := "DELETE FROM customer WHERE id=$1"

	smt, err := conn.Prepare(query)

	if err != nil {
		return
	}
	_, err = smt.Exec(id)
	return
}
