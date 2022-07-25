package db

import (
	"duguying/studio/modules/dbmodels"
	"time"

	"github.com/gogather/json"
	"gorm.io/gorm"
)

const (
	AGENT_STATUS_ALLOW  = 0
	AGENT_STATUS_FOBBID = 1

	AGENT_OFFLINE = 0
	AGENT_ONLINE  = 1
)

// CreateOrUpdateAgent 创建或更新 agent
func CreateOrUpdateAgent(tx *gorm.DB, clientId string, ip string) (agent *dbmodels.Agent, err error) {
	existAgent := &dbmodels.Agent{}
	err = tx.Table("agents").Where("client_id=?", clientId).First(existAgent).Error
	if err != nil {
		// not exist, create
		agent = &dbmodels.Agent{
			ClientId:    clientId,
			Ip:          ip,
			Online:      AGENT_ONLINE,
			Status:      AGENT_STATUS_ALLOW,
			OnlineTime:  time.Now(),
			OfflineTime: time.Now(),
		}
		err = tx.Table("agents").Create(agent).Error
		if err != nil {
			return nil, err
		}
	} else {
		// exist, update
		err = tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
			"online": AGENT_ONLINE,
			"status": AGENT_STATUS_ALLOW,
			"ip":     ip,
		}).Error
		if err != nil {
			return nil, err
		}
	}

	return agent, nil
}

func PutPerf(tx *gorm.DB, clientId string, os string, arch string, hostname string, ipIns []string) (err error) {
	ipInBytes, _ := json.Marshal(ipIns)

	err = tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
		"online":   AGENT_ONLINE,
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

// GetAgentByClientId 通过 clientId 获取 agent
func GetAgentByClientId(tx *gorm.DB, clientId string) (agent *dbmodels.Agent, err error) {
	agent = &dbmodels.Agent{}
	err = tx.Table("agents").Where("client_id=?", clientId).First(agent).Error
	if err != nil {
		return nil, err
	} else {
		return agent, nil
	}
}

// ListAllAvailableAgents 列出所有可用 agent
func ListAllAvailableAgents(tx *gorm.DB) (agents []*dbmodels.Agent, err error) {
	agents = []*dbmodels.Agent{}
	err = tx.Table("agents").Where("status=?", AGENT_STATUS_ALLOW).Find(&agents).Error
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

	err = tx.Table("agents").Where("id=?", id).Update("status", AGENT_STATUS_FOBBID).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateAgentOffline(tx *gorm.DB, clientId string) (err error) {
	err = tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
		"online":       AGENT_OFFLINE,
		"offline_time": time.Now(),
	}).Error
	if err != nil {
		return err
	}

	return nil
}
