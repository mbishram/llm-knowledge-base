package exercise

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Exercise struct {
	Activity  string    `json:"activity"`
	Duration  int       `json:"duration_minutes"`
	Timestamp time.Time `json:"timestamp"`
	Notes     string    `json:"notes"`
}

type Manager struct {
	DataDir string
}

func NewManager(dataDir string) (*Manager, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}
	return &Manager{DataDir: dataDir}, nil
}

func (m *Manager) LogExercise(activity string, duration int, notes string) (*Exercise, error) {
	ex := &Exercise{
		Activity:  activity,
		Duration:  duration,
		Timestamp: time.Now(),
		Notes:     notes,
	}

	// We append to a history file or save as individual files?
	// The plan says "Local file-based storage (JSON)".
	// Let's use a single file for history to make it easier to view history.
	historyFile := filepath.Join(m.DataDir, "exercise_history.json")
	
	var history []Exercise
	if data, err := os.ReadFile(historyFile); err == nil {
		_ = json.Unmarshal(data, &history)
	}

	history = append(history, *ex)

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(historyFile, data, 0644); err != nil {
		return nil, err
	}

	return ex, nil
}

func (m *Manager) GetHistory() ([]Exercise, error) {
	historyFile := filepath.Join(m.DataDir, "exercise_history.json")
	data, err := os.ReadFile(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Exercise{}, nil
		}
		return nil, err
	}

	var history []Exercise
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return history, nil
}
