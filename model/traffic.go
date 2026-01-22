package service

import (
	"time"
	"github.com/你的用户名/nezha-traffic/model"
	"gorm.io/gorm"
)

func UpdateTrafficUsage(db *gorm.DB, serverID uint, deltaRx, deltaTx int64) (*model.TrafficUsage, error) {
	now := time.Now()
	var usage model.TrafficUsage
	err := db.FirstOrCreate(&usage, model.TrafficUsage{
		ServerID: serverID,
		Year:     now.Year(),
		Month:    int(now.Month()),
	}).Error
	if err != nil {
		return nil, err
	}

	if deltaRx > 0 {
		usage.Rx += deltaRx
	}
	if deltaTx > 0 {
		usage.Tx += deltaTx
	}

	err = db.Save(&usage).Error
	return &usage, err
}
