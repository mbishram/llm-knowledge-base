package notes

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
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

func (m *Manager) AddNote(title, content string) (*Note, error) {
	note := &Note{
		ID:      uuid.New().String(),
		Title:   title,
		Content: content,
		Created: time.Now(),
	}

	filename := filepath.Join(m.DataDir, fmt.Sprintf("note_%s.json", note.ID))
	data, err := json.MarshalIndent(note, "", "  ")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return nil, err
	}

	return note, nil
}

func (m *Manager) GetNote(id string) (*Note, error) {
	filename := filepath.Join(m.DataDir, fmt.Sprintf("note_%s.json", id))
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var note Note
	if err := json.Unmarshal(data, &note); err != nil {
		return nil, err
	}

	return &note, nil
}

func (m *Manager) ListNotes() ([]Note, error) {
	files, err := filepath.Glob(filepath.Join(m.DataDir, "note_*.json"))
	if err != nil {
		return nil, err
	}

	var notes []Note
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		var note Note
		if err := json.Unmarshal(data, &note); err == nil {
			notes = append(notes, note)
		}
	}

	return notes, nil
}
