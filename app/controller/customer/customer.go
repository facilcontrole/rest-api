package customer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/facilcontrole/rest-api/app/middleware"
	"github.com/facilcontrole/rest-api/database/postgres"
	"github.com/facilcontrole/rest-api/logger"
)

type CustomerHandler struct {
	Repository Repository
}

func NewCustomerHandler(m *pat.PatternServeMux, rp Repository) {

	handler := &CustomerHandler{
		Repository: rp,
	}

	m.Get("/customer", middleware.Auth(http.HandlerFunc(handler.FindAll)))
	m.Post("/customer", middleware.Auth(http.HandlerFunc(handler.Storage)))

}

func (d *CustomerHandler) FindAll(w http.ResponseWriter, req *http.Request) {

	conn := postgres.App()
	defer conn.Close()
	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)
	lg.Conn = conn

	items, err := d.Repository.FindAll(conn)

	if err != nil {
		lg.Error(w, err)
		return
	}

	body, _ := json.Marshal(&items)

	lg.Create()

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

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
