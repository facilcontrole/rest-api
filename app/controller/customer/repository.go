package customer

import "database/sql"

type Repository interface {
	FindAll(conn *sql.DB) (items map[string]interface{}, err error)
}
