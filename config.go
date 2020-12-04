package main

import (
	"github.com/pelletier/go-toml"
)

type LoginInfo struct {
	UserName string `toml:"userName"`
	Password string `toml:"password"`
}

func GetLoginInfo() (*LoginInfo, error) {
	root, err := toml.LoadFile("config.toml")
	if err != nil {
		return nil, err
	}
	var loginInfo LoginInfo
	err = root.Unmarshal(&loginInfo)
	if err != nil {
		return nil, err
	}
	return &loginInfo, err
}
