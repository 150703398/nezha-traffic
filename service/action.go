package service

import (
	"fmt"
	"os/exec"

	"github.com/nezhahq/nezha/model"
)

// ExecuteAutoAction 执行超限动作
func ExecuteAutoAction(server model.Server) {
	if server.TrafficPlan == nil {
		return
	}

	switch server.TrafficPlan.AutoAction {

	case "shutdown":
		// ⚠️ 远程关机命令，需要服务器允许 SSH 密钥免密登录
		cmd := fmt.Sprintf("ssh root@%s 'shutdown -h now'", server.Name)
		_ = exec.Command("bash", "-c", cmd).Run()
	}
}
