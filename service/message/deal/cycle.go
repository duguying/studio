package deal

import (
	"duguying/blog/modules/logger"
	"duguying/blog/service/message/pipe"
	"log"
)

func Start() {
	go func() {
		for {
			select {
			case msg := <-pipe.In:
				logger.L("ws").Println("[ws] 收到原始消息:", msg)
				err := DealWithMessage(msg)
				if err != nil {
					log.Println("[ws] pipe deal with message error:", err)
				}
			}
		}
		logger.L("ws").Println("[ws] ⚠️⚠️⚠️消息处服务已终止")
	}()
}
