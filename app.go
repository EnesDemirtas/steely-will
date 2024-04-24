package main

import (
	"context"
	"log"

	"github.com/EnesDemirtas/steely-will/core"
	"github.com/EnesDemirtas/steely-will/storage"
)

// App struct
type App struct {
	ctx context.Context
}


type Habits struct {
	Habits []core.Habit
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllHabits() Habits {
	s := storage.New()

	habits, err := s.LoadAll()
	if err != nil {
		log.Fatal("err")
	}

	return Habits{
		Habits: habits,
	}
}

func (a *App) CreateHabit(name string, goalDays int) {
	s := storage.New()

	habit := core.New(name, goalDays)

	if err := s.Store(*habit); err != nil {
		log.Fatal(err)
	}
}
