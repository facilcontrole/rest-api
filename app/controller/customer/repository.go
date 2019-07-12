package customer

import "database/sql"

type Repository interface {
	FindAll(conn *sql.DB) (item map[string]interface{}, err error)
}
