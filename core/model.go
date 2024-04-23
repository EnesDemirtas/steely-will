package core

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID 				uuid.UUID `json:"id"`
	Name 			string	  `json:"name"`
	FailStack 		int		  `json:"failStack"`
	GoalDays		int 	  `json:"goalDays"`
	LastResetDate	time.Time `json:"lastResetDate"`
}