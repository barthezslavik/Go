package main

import (
	"fmt"
	"time"
)

// Equipment represents a piece of equipment being monitored
type Equipment struct {
	ID        string
	Model     string
	Operating int
	Failures  int
	LastCheck time.Time
}

// Check checks the equipment for failures
func (e *Equipment) Check() error {
	// Check for failures
	if e.Failures > 0 {
		return fmt.Errorf("equipment has failed")
	}
	return nil
}

// Maintenance schedules maintenance for the equipment
func (e *Equipment) Maintenance() {
	// Schedule maintenance
	e.LastCheck = time.Now()
}

func main() {
	// Create a new piece of equipment
	equipment := &Equipment{
		ID:        "123",
		Model:     "XYZ",
		Operating: 0,
		Failures:  0,
		LastCheck: time.Now(),
	}

	// Check the equipment
	if err := equipment.Check(); err != nil {
		fmt.Println("Error checking equipment:", err)
		// Schedule maintenance
		equipment.Maintenance()
	}
}
