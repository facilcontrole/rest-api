package routes

import (
	"github.com/bmizerany/pat"
	"github.com/facilcontrole/app/controller/customer"
	crRepo "github.com/facilcontrole/app/controller/customer/repository"
)

func Routes() (m *pat.PatternServeMux) {

	m = pat.New()

	rp := crRepo.NewPostgresRepository()
	customer.NewCustomerHandler(m, rp)

	return

}
