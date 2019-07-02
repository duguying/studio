package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/models"
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
func CreateOrUpdateAgent(clientId string, ip string) (agent *models.Agent, err error) {
	tx := g.Db.Begin()
	existAgent := &models.Agent{}
	errs := tx.Table("agents").Where("client_id=?", clientId).First(existAgent).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		// not exist, create
		agent = &models.Agent{
			ClientId:    clientId,
			Ip:          ip,
			Online:      AGENT_ONLINE,
			Status:      AGENT_STATUS_ALLOW,
			OnlineTime:  time.Now(),
			OfflineTime: time.Now(),
		}
		errs = tx.Table("agents").Create(agent).GetErrors()
		if len(errs) > 0 && errs[0] != nil {
			tx.Rollback()
			return nil, errs[0]
		}
	} else {
		// exist, update
		errs = tx.Table("agents").Updates(map[string]interface{}{
			"online": AGENT_ONLINE,
			"ip":     ip,
		}).Where("client_id=?", clientId).GetErrors()
		if len(errs) > 0 && errs[0] != nil {
			tx.Rollback()
			return nil, errs[0]
		}
	}

	errs = tx.Commit().GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	} else {
		return agent, nil
	}
}

func PutPerf(clientId string, os string, arch string, hostname string, ipIns []string) (err error) {
	tx := g.Db.Begin()
	ipInBytes, _ := json.Marshal(ipIns)

	errs := tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
		"online":   AGENT_ONLINE,
		"os":       os,
		"arch":     arch,
		"hostname": hostname,
		"ip_ins":   string(ipInBytes),
	}).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		tx.Rollback()
		return errs[0]
	}

	errs = tx.Commit().GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	} else {
		return nil
	}
}

// 通过 id 获取 agent
func GetAgent(id uint) (agent *models.Agent, err error) {
	agent = &models.Agent{}
	errs := g.Db.Table("agents").Where("id=?", id).First(agent).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	} else {
		return agent, nil
	}
}

// 通过 clientId 获取 agent
func GetAgentByClientId(clientId string) (agent *models.Agent, err error) {
	agent = &models.Agent{}
	errs := g.Db.Table("agents").Where("client_id=?", clientId).First(agent).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	} else {
		return agent, nil
	}
}

// 列出所有可用 agent
func ListAllAvailableAgents() (agents []*models.Agent, err error) {
	agents = []*models.Agent{}
	errs := g.Db.Table("agents").Where("status=?", AGENT_STATUS_ALLOW).Find(&agents).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	} else {
		return agents, nil
	}
}

// 禁用 agent
func ForbidAgent(id uint) (err error) {
	agent := &models.Agent{}
	errs := g.Db.Table("agents").Where("id=?", id).First(agent).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}

	errs = g.Db.Table("agents").Where("id=?", id).Update("status", AGENT_STATUS_FOBBID).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}

	return nil
}

func UpdateAgentOffline(clientId string) (err error) {
	tx := g.Db.Begin()

	errs := tx.Table("agents").Where("client_id=?", clientId).Updates(map[string]interface{}{
		"online":       AGENT_OFFLINE,
		"offline_time": time.Now(),
	}).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		tx.Rollback()
		return errs[0]
	}

	errs = tx.Commit().GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	} else {
		return nil
	}
}
