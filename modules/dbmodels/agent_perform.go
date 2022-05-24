package dbmodels

import "github.com/gogather/json"

type AgentPerform struct {
	Id        uint   `json:"id"`
	Timestamp uint64 `json:"timestamp"`
	ClientId  string `json:"client_id"`
}

func (ap *AgentPerform) String() string {
	c, _ := json.Marshal(ap)
	return string(c)
}
