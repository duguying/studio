package models

import (
	"encoding/json"
	"time"
)

// APILog api日志
type APILog struct {
	ID        uint      `json:"id,omitempty"`
	Method    string    `json:"method"`
	URI       string    `json:"uri"`
	Query     string    `json:"query"`
	Body      string    `json:"body"`
	Ok        bool      `json:"ok"`
	Response  string    `json:"response"`
	ClientIP  string    `json:"client_ip"`
	RequestID string    `json:"request_id"`
	Cost      string    `json:"cost"`
	CreatedAt time.Time `json:"created_at"`
}

func (al *APILog) ToMap() map[string]interface{} {
	obj := map[string]interface{}{}
	c, _ := json.Marshal(al)
	_ = json.Unmarshal(c, &obj)
	return obj
}
