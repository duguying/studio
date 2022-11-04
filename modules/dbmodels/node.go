package dbmodels

import "time"

// Node 服务节点
type Node struct {
	UUID

	NodeID       string    `json:"node_id"`
	AccessIPPort string    `json:"access_ipport"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
