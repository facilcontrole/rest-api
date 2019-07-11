package models

type Logger struct {
	ID     string     `json:"id"`
	Http   LoggerHttp `json:"http"`
	UserID string     `json:"user_id"`
}

type LoggerHttp struct {
	RemoteAddr string `json:"remote_addr"`
	Action     string `json:"action"`
	Method     string `json:"method"`
	RawQuery   string `json:"raw_query"`
	UserAgent  string `json:"user_agent"`
	Status     int    `json:"status"`
	Error      string `json:"error"`
	Body       string `json:"body"`
}
