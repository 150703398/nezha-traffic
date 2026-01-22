package model

type Server struct {
	ID            uint
	Name          string
	TrafficPlanID *uint
	TrafficPlan   *TrafficPlan
}
