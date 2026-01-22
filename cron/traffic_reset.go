package cron

import (
	"time"
	"github.com/你的用户名/nezha-traffic/dao"
	"github.com/你的用户名/nezha-traffic/model"
)

func MonthlyReset() {
	now := time.Now()
	if now.Day() != 1 {
		return
	}

	dao.DB.Where("year != ? OR month != ?", now.Year(), int(now.Month())).Delete(&model.TrafficUsage{})
}
