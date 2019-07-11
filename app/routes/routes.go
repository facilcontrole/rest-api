package routes

import (
	"github.com/bmizerany/pat"
	"github.com/facilcontrole/rest-api/app/controller/customer"
	"github.com/facilcontrole/rest-api/database/postgres"
)

func Routes() (m *pat.PatternServeMux) {

	m = pat.New()

	conn := postgres.App()
	defer conn.Close()

	customer.NewCustomerHandler(m, conn)

	return

}
