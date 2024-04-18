package core

import "time"

type Habit struct {
	Name 			string
	FailStack 		int
	GoalDate		time.Time
	LastResetDate	time.Time
}