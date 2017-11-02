// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package models

type Users struct {
	Id       int
	Username string
	Password string
	Salt     string
	Email    string
}
