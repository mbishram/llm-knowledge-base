package exercise

import (
	"os"
	"testing"
)

func TestManager_LogExercise(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exercise_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	mgr, err := NewManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	activity := "Running"
	duration := 30
	notes := "Morning run"
	ex, err := mgr.LogExercise(activity, duration, notes)
	if err != nil {
		t.Fatalf("Failed to log exercise: %v", err)
	}

	if ex.Activity != activity {
		t.Errorf("Expected activity %s, got %s", activity, ex.Activity)
	}

	if ex.Duration != duration {
		t.Errorf("Expected duration %d, got %d", duration, ex.Duration)
	}

	history, err := mgr.GetHistory()
	if err != nil {
		t.Fatalf("Failed to get history: %v", err)
	}

	if len(history) != 1 {
		t.Errorf("Expected history length 1, got %d", len(history))
	}

	if history[0].Activity != activity {
		t.Errorf("Expected history activity %s, got %s", activity, history[0].Activity)
	}
}
