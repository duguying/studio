// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/29.

package store

import (
	"testing"
	"time"
	"fmt"
)

func TestAgentStatusInfo_String(t *testing.T) {
	fmt.Println(time.Now().Add(-time.Hour*24*365).Format(time.RFC3339))
}
