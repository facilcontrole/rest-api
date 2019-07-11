package logger

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/facilcontrole/rest-api/app/models"
)

type Logger struct {
	Items models.Logger `json:"items"`
	Conn  *sql.DB       `json:"conn,omitempty"`
}

func (data Logger) Init(UserId string, status int, req *http.Request) (item models.Logger) {

	item.UserID = UserId

	item.Http.UserAgent = req.UserAgent()
	item.Http.RemoteAddr = req.RemoteAddr
	item.Http.RawQuery = req.URL.RawQuery

	item.Http.Method = req.Method
	item.Http.Action = req.URL.Path
	item.Http.Status = status

	return

}

func (data Logger) Create() {

	body, _ := json.Marshal(&data)

	println(string(body))

	return

}

func (lg Logger) Error(w http.ResponseWriter, err error) {

	lg.Items.Http.Error = err.Error()
	lg.Items.Http.Status = http.StatusBadRequest
	lg.Create()

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))

}
