package pipe

import (
	"github.com/gogather/safemap"
	"log"
	"duguying/blog/service/message/model"
)

type ClientPipe struct {
	clientId string
	out      chan model.Msg
}

var In chan model.Msg
var pm *safemap.SafeMap // [clientId] -> ClientPipe
var cm *safemap.SafeMap // [clientId] -> connId

func InitPipeline() {
	In = make(chan model.Msg, 100)
	pm = safemap.New()
	cm = safemap.New()
	conns = safemap.New()
}

func AddUserPipe(clientId string, out chan model.Msg, connId string) {
	log.Printf("注册设备 ID:%s\n", clientId)
	pm.Put(clientId, &ClientPipe{
		clientId: clientId,
		out:      out,
	})
	cm.Put(clientId, connId)
}

func RemoveUserPipe(clientId string) {
	pm.Remove(clientId)
	cm.Remove(clientId)
}

func SendMsg(clientId string, msg model.Msg) (success bool, err error) {
	iCli, exist := pm.Get(clientId)
	if !exist {
		return false, nil
	}
	cli := iCli.(*ClientPipe)
	cli.out <- msg
	return true, nil
}

func GetConMap() *safemap.SafeMap {
	return cm
}