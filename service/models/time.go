// Copyright 2019. All rights reserved.
// This file is part of sparta-admin project
// I am coding in Tencent
// Created by rainesli on 2019/3/18.

package models

import (
	"github.com/json-iterator/go"
	"time"
	"unsafe"
)

func RegisterTimeAsLayoutCodec(layout string) {
	jsoniter.RegisterTypeEncoder("time.Time", &timeAsString{layout: layout})
	jsoniter.RegisterTypeDecoder("time.Time", &timeAsString{layout: layout})
}

type timeAsString struct {
	layout string
}

func (codec *timeAsString) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	tm, err := time.Parse(codec.layout, iter.ReadString())
	if err != nil {
		return
	}
	*((*time.Time)(ptr)) = tm
}

func (codec *timeAsString) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.UnixNano() == 0
}

func (codec *timeAsString) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	op := ts.Format(codec.layout)
	stream.WriteString(op)
}
