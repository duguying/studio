package deal

import (
	"duguying/studio/modules/logger"
	"duguying/studio/service/message/model"
	"duguying/studio/service/message/pipe"
	"duguying/studio/service/message/store"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

func DealWithMessage(rcvMsgPack model.Msg) (err error) {
	data := rcvMsgPack.Data
	switch rcvMsgPack.Cmd {
	case model.CMD_PERF:
		{
			err := store.PutPerf(rcvMsgPack.ClientId, uint64(time.Now().Unix()), data)
			if err != nil {
				log.Println("boltdb store data failed, err:", err.Error())
			}
			return nil
		}
	case model.CMD_CLI_PIPE:
		{
			pipeData := model.CliPipe{}
			err := proto.Unmarshal(data, &pipeData)
			if err != nil {
				log.Println("parse pipe data failed, err:", err.Error())
				return err
			}
			session := pipeData.Session
			pid := pipeData.Pid
			pair, exist := pipe.GetCliChanPair(session, pid)
			if exist {
				//log.Println("cli ---> xterm:", pipeData.Data)
				logger.L("agentsnt").Printf("agent receive: %s\n", string(pipeData.Data))
				pair.ChanIn <- append([]byte{model.TERM_PIPE}, pipeData.Data...)
			}
			return nil
		}
	case model.CMD_CLI_CMD:
		{
			pcmd := model.CliCmd{}
			err := proto.Unmarshal(data, &pcmd)
			if err != nil {
				log.Println("parse pipe cmd data failed, err:", err.Error())
				return err
			}
			if pcmd.Cmd == model.CliCmd_OPEN {
				pipe.SetCliPid(pcmd.Session, pcmd.RequestId, pcmd.Pid)
			} else if pcmd.Cmd == model.CliCmd_CLOSE {
				// 1. close ws
				ws, exist := pipe.GetPidCon(pcmd.Session, pcmd.Pid)
				if exist {
					ws.Close()
					pipe.DelPidCon(pcmd.Session, pcmd.Pid)
				}

				// 2. clear key
				pipe.DelCliPid(pcmd.Session, pcmd.RequestId)
			}
			return nil
		}
	}
	return nil
}
