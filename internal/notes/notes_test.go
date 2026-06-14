package notes

import (
	"os"
	"path/filepath"
	"testing"
)

func TestManager_AddNote(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "notes_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	mgr, err := NewManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	title := "Test Note"
	content := "This is a test note content."
	note, err := mgr.AddNote(title, content)
	if err != nil {
		t.Fatalf("Failed to add note: %v", err)
	}

	if note.Title != title {
		t.Errorf("Expected title %s, got %s", title, note.Title)
	}

	if note.Content != content {
		t.Errorf("Expected content %s, got %s", content, note.Content)
	}

	// Verify file exists
	filename := filepath.Join(tempDir, "note_"+note.ID+".json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist", filename)
	}
}

func TestManager_ListNotes(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "notes_list_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	mgr, err := NewManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	mgr.AddNote("Note 1", "Content 1")
	mgr.AddNote("Note 2", "Content 2")

	notes, err := mgr.ListNotes()
	if err != nil {
		t.Fatalf("Failed to list notes: %v", err)
	}

	if len(notes) != 2 {
		t.Errorf("Expected 2 notes, got %d", len(notes))
	}
}
