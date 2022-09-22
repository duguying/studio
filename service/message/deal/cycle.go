package deal

import (
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
	}()
}
