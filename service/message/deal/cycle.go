package deal

import (
	"duguying/studio/modules/logger"
	"duguying/studio/service/message/pipe"
	"log"
)

func Start() {
	go func() {
		for {
			select {
			case msg := <-pipe.In:
				err := DealWithMessage(msg)
				if err != nil {
					log.Println("[ws] pipe deal with message error:", err)
				}
			}
		}
		logger.L("ws").Println("[ws] ⚠️⚠️⚠️消息处服务已终止")
	}()
}
