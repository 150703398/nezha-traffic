package model

type TrafficUsage struct {
	ServerID uint `gorm:"primaryKey"`
	Year     int  `gorm:"primaryKey"`
	Month    int  `gorm:"primaryKey"`
	Rx       int64
	Tx       int64
	Exceeded bool
}
