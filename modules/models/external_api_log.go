package models

import (
	"time"

	"github.com/gogather/json"
)

type ExternalApiLog struct {
	Id        uint      `json:"id"`
	Method    string    `json:"method"`
	Uri       string    `json:"uri"`
	Query     string    `json:"query"`
	Body      string    `json:"body"`
	Ok        bool      `json:"ok"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *ExternalApiLog) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}
