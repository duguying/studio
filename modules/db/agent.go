package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"time"

	"github.com/gogather/json"
)

const (
	AGENT_STATUS_ALLOW  = 0
	AGENT_STATUS_FOBBID = 1

	AGENT_OFFLINE = 0
	AGENT_ONLINE  = 1
)

// 创建或更新 agent
func CreateOrUpdateAgent(clientId string, ip string) (agent *dbmodels.Agent, err error) {
	tx := g.Db.Begin()
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
			tx.Rollback()
			return nil, err
		}
	} else {
		// exist, update
		err = tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
			"online": AGENT_ONLINE,
			"ip":     ip,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	} else {
		return agent, nil
	}
}

func PutPerf(clientId string, os string, arch string, hostname string, ipIns []string) (err error) {
	tx := g.Db.Begin()
	ipInBytes, _ := json.Marshal(ipIns)

	err = tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
		"online":   AGENT_ONLINE,
		"os":       os,
		"arch":     arch,
		"hostname": hostname,
		"ip_ins":   string(ipInBytes),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 通过 id 获取 agent
func GetAgent(id uint) (agent *dbmodels.Agent, err error) {
	agent = &dbmodels.Agent{}
	err = g.Db.Table("agents").Where("id=?", id).First(agent).Error
	if err != nil {
		return nil, err
	} else {
		return agent, nil
	}
}

// 通过 clientId 获取 agent
func GetAgentByClientId(clientId string) (agent *dbmodels.Agent, err error) {
	agent = &dbmodels.Agent{}
	err = g.Db.Table("agents").Where("client_id=?", clientId).First(agent).Error
	if err != nil {
		return nil, err
	} else {
		return agent, nil
	}
}

// 列出所有可用 agent
func ListAllAvailableAgents() (agents []*dbmodels.Agent, err error) {
	agents = []*dbmodels.Agent{}
	err = g.Db.Table("agents").Where("status=?", AGENT_STATUS_ALLOW).Find(&agents).Error
	if err != nil {
		return nil, err
	} else {
		return agents, nil
	}
}

// 禁用 agent
func ForbidAgent(id uint) (err error) {
	agent := &dbmodels.Agent{}
	err = g.Db.Table("agents").Where("id=?", id).First(agent).Error
	if err != nil {
		return err
	}

	err = g.Db.Table("agents").Where("id=?", id).Update("status", AGENT_STATUS_FOBBID).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateAgentOffline(clientId string) (err error) {
	tx := g.Db.Begin()

	err = tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
		"online":       AGENT_OFFLINE,
		"offline_time": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
