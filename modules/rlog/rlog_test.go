package rlog

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestRLog(t *testing.T) {
	id := uuid.New().String()
	rl, err := NewEsAdaptor("http://jump.duguying.net:19200", "test")
	if err != nil {
		fmt.Println(err)
	}
	rl.Report(map[string]interface{}{
		"name":  "rex",
		"age":   32,
		"phone": 123456,
		"uuid":  id,
	})
}
