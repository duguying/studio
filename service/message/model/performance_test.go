// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/7.

package model

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestPerformanceMonitor(t *testing.T) {
	perf := &PerformanceMonitor{
		Mem: &PerformanceMonitor_Memory{
			TotalMem:   1024,
			UsedMem:    824,
			FreeMem:    200,
			ActualUsed: 100,
			ActualFree: 200,
			TotalSwap:  100,
			UsedSwap:   20,
			FreeSwap:   80,
		},
	}
	data, err := proto.Marshal(perf)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
	newTest := &PerformanceMonitor{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if perf.GetMem() != newTest.GetMem() {
		log.Fatalf("data mismatch %q != %q", perf.GetMem(), newTest.GetMem())
	}
}
