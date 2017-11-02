// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package g

import (
	"duguying/blog/modules/configuration"
	"github.com/jinzhu/gorm"
)

var (
	Config *configuration.Config
	Db *gorm.DB

	InstallMode bool = false

	GitVersion = "00000000"
	BuildTime  = "2000-01-01T00:00:00+0800"
)
