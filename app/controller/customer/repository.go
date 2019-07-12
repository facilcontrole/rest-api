package customer

import (
	"database/sql"

	"github.com/facilcontrole/rest-api/app/models"
)

type Repository interface {
	FindAll(conn *sql.DB) (item map[string]interface{}, err error)
	Storage(name string, body []byte, conn *sql.DB) (id string, err error)
	FindById(id string, conn *sql.DB) (items, name string, err error)
	Update(item models.Customer, conn *sql.DB) (err error)
	Delete(id string, conn *sql.DB) (err error)
}
