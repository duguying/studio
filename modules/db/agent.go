package db

import (
	"duguying/studio/modules/dbmodels"
	"time"

	"github.com/gogather/json"
	"gorm.io/gorm"
)

const (
	AgentStatusAllow   = 0
	AgentStatusForbbid = 1

	AgentOffline = 0
	AgentOnline  = 1
)

// CreateOrUpdateAgent 创建或更新 agent
func CreateOrUpdateAgent(tx *gorm.DB, clientID string, IP string) (agent *dbmodels.Agent, err error) {
	existAgent := &dbmodels.Agent{}
	err = tx.Table("agents").Where("client_id=?", clientID).First(existAgent).Error
	if err != nil {
		// not exist, create
		agent = &dbmodels.Agent{
			ClientID:    clientID,
			IP:          IP,
			Online:      AgentOnline,
			Status:      AgentStatusAllow,
			OnlineTime:  time.Now(),
			OfflineTime: time.Now(),
		}
		err = tx.Table("agents").Create(agent).Error
		if err != nil {
			return nil, err
		}
	} else {
		// exist, update
		err = tx.Table("agents").Where("client_id=?", clientID).Updates(map[string]interface{}{
			"online": AgentOnline,
			"status": AgentStatusAllow,
			"ip":     IP,
		}).Error
		if err != nil {
			return nil, err
		}
	}

	return agent, nil
}

func PutPerf(tx *gorm.DB, clientID string, os string, arch string, hostname string, IPIns []string) (err error) {
	ipInBytes, _ := json.Marshal(IPIns)

	err = tx.Table("agents").Where("client_id=?", clientID).Updates(map[string]interface{}{
		"online":   AgentOnline,
		"os":       os,
		"arch":     arch,
		"hostname": hostname,
		"ip_ins":   string(ipInBytes),
	}).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAgent 通过 id 获取 agent
func GetAgent(tx *gorm.DB, id uint) (agent *dbmodels.Agent, err error) {
	agent = &dbmodels.Agent{}
	err = tx.Table("agents").Where("id=?", id).First(agent).Error
	if err != nil {
		return nil, err
	} else {
		return agent, nil
	}
}

// GetAgentByClientID 通过 clientID 获取 agent
func GetAgentByClientID(tx *gorm.DB, clientID string) (agent *dbmodels.Agent, err error) {
	agent = &dbmodels.Agent{}
	err = tx.Table("agents").Where("client_id=?", clientID).First(agent).Error
	if err != nil {
		return nil, err
	} else {
		return agent, nil
	}
}

// ListAllAvailableAgents 列出所有可用 agent
func ListAllAvailableAgents(tx *gorm.DB) (agents []*dbmodels.Agent, err error) {
	agents = []*dbmodels.Agent{}
	err = tx.Table("agents").Where("status=?", AgentStatusAllow).Find(&agents).Error
	if err != nil {
		return nil, err
	} else {
		return agents, nil
	}
}

// ForbidAgent 禁用 agent
func ForbidAgent(tx *gorm.DB, id uint) (err error) {
	agent := &dbmodels.Agent{}
	err = tx.Table("agents").Where("id=?", id).First(agent).Error
	if err != nil {
		return err
	}

	err = tx.Table("agents").Where("id=?", id).Update("status", AgentStatusForbbid).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateAgentOffline(tx *gorm.DB, clientID string) (err error) {
	err = tx.Table("agents").Where("client_id=?", clientID).Updates(map[string]interface{}{
		"online":       AgentOffline,
		"offline_time": time.Now(),
	}).Error
	if err != nil {
		return err
	}

	return nil
}
