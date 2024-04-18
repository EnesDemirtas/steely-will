package core

import "time"


func (h *Habit) New(name string, goalDate time.Time) *Habit {
	return &Habit{
		Name:			name,
		FailStack: 		0,
		GoalDate: 		goalDate,
		LastResetDate: 	time.Now(),
	}
}

func (h *Habit) GetTimeElapsed() time.Duration {
	return time.Since(h.LastResetDate)
}

func (h *Habit) Reset() {
	h.FailStack++
	h.LastResetDate = time.Now()
}

func (h *Habit) ChangeName(newName string) {
	h.Name = newName
}

func (h *Habit) ChangeGoalDate(newGoal time.Time) {
	h.GoalDate = newGoal
}

func (h *Habit) IsGoalReached() bool {
	return time.Now().After(h.GoalDate)
}