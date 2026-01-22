package controller

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/你的用户名/nezha-traffic/dao"
	"github.com/你的用户名/nezha-traffic/model"
)

func GetTraffic(c *gin.Context) {
	id := c.Param("id")
	now := time.Now()
	var usage model.TrafficUsage

	dao.DB.Where("server_id=? AND year=? AND month=?", id, now.Year(), int(now.Month())).First(&usage)

	var server model.Server
	dao.DB.Preload("TrafficPlan").First(&server, id)

	limit := int64(0)
	if server.TrafficPlan != nil {
		limit = server.TrafficPlan.MonthlyLimit
	}

	used := usage.Rx + usage.Tx
	percent := float64(0)
	if limit > 0 {
		percent = float64(used) / float64(limit) * 100
	}

	c.JSON(200, gin.H{
		"used":     used,
		"limit":    limit,
		"percent":  percent,
		"exceeded": usage.Exceeded,
	})
}
