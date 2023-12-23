package qweather

import (
	"go.dtapp.net/golog"
)

type Client struct {
	key     string
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

func NewClient(key string) (*Client, error) {
	return &Client{key: key}, nil
}
