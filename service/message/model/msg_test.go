package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsg(t *testing.T) {
	m := Msg{
		Type:     0,
		Cmd:      0,
		ClientId: "",
		Data:     nil,
	}

	assert.Equal(t, m.String(), `{"client_id":"","cmd":0,"data":"[]","type":0}`)
}
