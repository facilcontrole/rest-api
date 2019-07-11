package customer

import (
	"database/sql"
	"io/ioutil"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/facilcontrole/rest-api/app/middleware"
	"github.com/facilcontrole/rest-api/logger"
)

type CustomerHandler struct {
	Conn *sql.DB
}

func NewCustomerHandler(m *pat.PatternServeMux, conn *sql.DB) {

	handler := &CustomerHandler{
		Conn: conn,
	}

	m.Get("/customer", middleware.Auth(http.HandlerFunc(handler.FindAll)))
	m.Post("/customer", middleware.Auth(http.HandlerFunc(handler.Storage)))

}

func (d *CustomerHandler) FindAll(w http.ResponseWriter, req *http.Request) {

	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)
	lg.Conn = d.Conn

	lg.Create()

}

func (d *CustomerHandler) Storage(w http.ResponseWriter, req *http.Request) {

	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)

	_, err := ioutil.ReadAll(req.Body)

	if err != nil {
		lg.Error(w, err)
		return
	}

	lg.Create()

}
