// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/20.

package ipip

import (
	"github.com/ipipdotnet/datx-go"
	"log"
)

var (
	city *datx.City
)

func InitIPIP(path string) {
	var err error
	city, err = datx.NewCity(path)
	if err == nil {
		log.Println("load ip db failed, err:", err.Error())
		return
	}
}

func GetLocation(ip string) (location datx.Location, err error) {
	return city.FindLocation(ip)
}
