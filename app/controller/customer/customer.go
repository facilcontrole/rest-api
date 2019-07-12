package customer

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/facilcontrole/rest-api/app/middleware"
	"github.com/facilcontrole/rest-api/app/models"
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
	m.Get("/customer/:id", middleware.Auth(http.HandlerFunc(handler.FindById)))
	m.Put("/customer/:id", middleware.Auth(http.HandlerFunc(handler.Update)))
	m.Post("/customer", middleware.Auth(http.HandlerFunc(handler.Storage)))
	m.Del("/customer/:id", middleware.Auth(http.HandlerFunc(handler.Remove)))
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

	conn := postgres.App()
	defer conn.Close()

	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		lg.Error(w, err)
		return
	}

	var customer models.Customer

	err = json.Unmarshal(body, &customer)

	if err != nil {
		lg.Error(w, err)
		return
	}

	body, err = json.Marshal(&customer.Items)

	if err != nil {
		lg.Error(w, err)
		return
	}

	id, err := d.Repository.Storage(customer.Name, body, conn)

	if err != nil {
		lg.Error(w, err)
		return
	}

	body, err = d.findById(id, conn)

	if err != nil {
		lg.Error(w, err)
		return
	}
	lg.Create()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)

}

func (d *CustomerHandler) findById(id string, conn *sql.DB) (body []byte, err error) {

	items, name, err := d.Repository.FindById(id, conn)

	if err != nil {
		return
	}

	var customer models.Customer

	err = json.Unmarshal([]byte(items), &customer.Items)

	if err != nil {
		return
	}

	customer.ID = id
	customer.Name = name

	body, err = json.Marshal(&customer)
	return
}

func (d *CustomerHandler) FindById(w http.ResponseWriter, req *http.Request) {

	conn := postgres.App()
	defer conn.Close()

	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)
	lg.Conn = conn

	id := req.URL.Query().Get(":id")

	body, err := d.findById(id, conn)
	if err != nil {
		lg.Error(w, err)
		return
	}
	lg.Create()

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func (d *CustomerHandler) Update(w http.ResponseWriter, req *http.Request) {

	conn := postgres.App()
	defer conn.Close()

	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)
	lg.Conn = conn

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		lg.Error(w, err)
		return
	}

	id := req.URL.Query().Get(":id")

	var customer models.Customer

	err = json.Unmarshal(body, &customer)

	if err != nil {
		lg.Error(w, err)
		return
	}
	customer.ID = id
	err = d.Repository.Update(customer, conn)

	if err != nil {
		lg.Error(w, err)
		return
	}

	body, err = d.findById(id, conn)

	if err != nil {
		lg.Error(w, err)
		return
	}
	lg.Create()

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func (d *CustomerHandler) Remove(w http.ResponseWriter, req *http.Request) {

	conn := postgres.App()
	defer conn.Close()

	var lg logger.Logger
	lg.Items = lg.Init(middleware.UserId, http.StatusOK, req)
	lg.Conn = conn

	id := req.URL.Query().Get(":id")

	err := d.Repository.Delete(id, conn)

	if err != nil {
		lg.Error(w, err)
		return
	}
	lg.Create()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
