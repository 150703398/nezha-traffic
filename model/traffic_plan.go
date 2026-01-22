package model

type TrafficPlan struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string
	MonthlyLimit int64
	AutoAction   string // none
	Notify       bool
}
