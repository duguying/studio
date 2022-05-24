// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/20.

package ipip

import (
	"github.com/ipipdotnet/ipdb-go"
	"log"
)

var (
	db *ipdb.City
)

func InitIPIP(path string) {
	var err error
	db, err = ipdb.NewCity(path)
	if err != nil {
		log.Fatal(err)
	}
}

func GetLocation(ip string) (location *ipdb.CityInfo, err error) {
	return db.FindInfo(ip,"CN")
}
