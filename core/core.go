package core

import (
	"time"

	"github.com/google/uuid"
)


func New(name string, goalDays int) *Habit {
	return &Habit{
		ID:				uuid.New(),
		Name:			name,
		FailStack: 		0,
		GoalDays: 		goalDays,
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

func (h *Habit) ChangeGoalDays(newGoal int) {
	h.GoalDays = newGoal
}

func (h *Habit) IsGoalReached() bool {
	return time.Now().After(h.LastResetDate.Add(time.Duration(h.GoalDays) * 24 * time.Hour))
}