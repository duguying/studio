package deal

import (
	"duguying/blog/service/message/model"
	"duguying/blog/service/message/store"
	"log"
	"time"
)

func DealWithMessage(rcvMsgPack model.Msg) (err error) {
	switch rcvMsgPack.Cmd {
	case model.CMD_PERF:
		{
			err := store.Put(rcvMsgPack.ClientId, uint64(time.Now().Unix()), rcvMsgPack.Data)
			if err != nil {
				log.Println("boltdb store data failed, err:", err.Error())
			}
			return nil
		}
	}
	return nil
}
