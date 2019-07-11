package customer

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/facilcontrole/rest-api/app/middleware"
)

type CustomerHandler struct {
}

func NewCustomerHandler(m *pat.PatternServeMux) {

	handler := &CustomerHandler{}

	m.Get("/customer", middleware.Auth(http.HandlerFunc(handler.FindAll)))
	m.Post("/customer", middleware.Auth(http.HandlerFunc(handler.Storage)))

}

func (d *CustomerHandler) FindAll(w http.ResponseWriter, req *http.Request) {

}

func (d *CustomerHandler) Storage(w http.ResponseWriter, req *http.Request) {

}
