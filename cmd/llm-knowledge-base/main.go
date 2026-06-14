package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mbishram/llm-knowledge-base/internal/exercise"
	"github.com/mbishram/llm-knowledge-base/internal/git"
	"github.com/mbishram/llm-knowledge-base/internal/notes"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	dataDir := "data"
	notesDir := filepath.Join(dataDir, "notes")
	exerciseDir := filepath.Join(dataDir, "exercises")

	if err := os.MkdirAll(notesDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create notes dir: %v\n", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(exerciseDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create exercises dir: %v\n", err)
		os.Exit(1)
	}

	notesMgr, err := notes.NewManager(notesDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init notes manager: %v\n", err)
		os.Exit(1)
	}

	exerciseMgr, err := exercise.NewManager(exerciseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init exercise manager: %v\n", err)
		os.Exit(1)
	}

	// Initialize the MCP server
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "llm-knowledge-base",
			Version: "1.0.0",
		},
		nil,
	)

	// Add a simple ping tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ping",
		Description: "A simple tool to verify connectivity",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, string, error) {
		return nil, "pong", nil
	})

	// add_note tool
	type AddNoteArgs struct {
		Title   string `json:"title" mcp:"The title of the note"`
		Content string `json:"content" mcp:"The content of the note"`
	}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "add_note",
		Description: "Create a new note",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args AddNoteArgs) (*mcp.CallToolResult, string, error) {
		note, err := notesMgr.AddNote(args.Title, args.Content)
		if err != nil {
			return nil, "", err
		}

		// Git commit
		_ = git.Commit(notesDir, fmt.Sprintf("Added note: %s", note.Title))

		return nil, fmt.Sprintf("Note created with ID: %s", note.ID), nil
	})

	// list_notes tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_notes",
		Description: "List all notes",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, []notes.Note, error) {
		list, err := notesMgr.ListNotes()
		if err != nil {
			return nil, nil, err
		}
		return nil, list, nil
	})

	// get_note tool
	type GetNoteArgs struct {
		ID string `json:"id" mcp:"The ID of the note to read"`
	}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_note",
		Description: "Read a specific note by ID",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args GetNoteArgs) (*mcp.CallToolResult, *notes.Note, error) {
		note, err := notesMgr.GetNote(args.ID)
		if err != nil {
			return nil, nil, err
		}
		return nil, note, nil
	})

	// log_exercise tool
	type LogExerciseArgs struct {
		Activity string `json:"activity" mcp:"The exercise activity (e.g., 'Running')"`
		Duration int    `json:"duration" mcp:"The duration in minutes"`
		Notes    string `json:"notes" mcp:"Optional notes about the exercise"`
	}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "log_exercise",
		Description: "Log an exercise activity",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args LogExerciseArgs) (*mcp.CallToolResult, string, error) {
		ex, err := exerciseMgr.LogExercise(args.Activity, args.Duration, args.Notes)
		if err != nil {
			return nil, "", err
		}

		// Git commit
		_ = git.Commit(exerciseDir, fmt.Sprintf("Logged exercise: %s", ex.Activity))

		return nil, fmt.Sprintf("Logged %d minutes of %s", ex.Duration, ex.Activity), nil
	})

	// get_exercise_history tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_exercise_history",
		Description: "Get the exercise history",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, []exercise.Exercise, error) {
		history, err := exerciseMgr.GetHistory()
		if err != nil {
			return nil, nil, err
		}
		return nil, history, nil
	})

	// Run the server on stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
