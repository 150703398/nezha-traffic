package service

import (
	"fmt"
	"github.com/你的用户名/nezha-traffic/model"
)

func NotifyTrafficExceeded(server model.Server, usage model.TrafficUsage) {
	msg := fmt.Sprintf(
		"🚨 流量超限\n服务器：%s\n已用：%.2f GB",
		server.Name,
		float64(usage.Rx+usage.Tx)/1024/1024/1024,
	)
	SendTelegram(msg)
}

func SendTelegram(msg string) {
	// 调用已有 telegram 通知逻辑
}
