# LLM Knowledge Base

A modular system for managing personal knowledge (notes) and exercise logs via the Model Context Protocol (MCP).

## Project Structure

- `agents/`: Contains agent configuration files like `AGENT.md`.
- `bin/`: Compiled binaries.
- `cmd/`: Entry points for the application.
    - `llm-knowledge-base/`: The main MCP server.
- `data/`: Local file-based storage (JSON).
    - `notes/`: Individual note files.
    - `exercises/`: Exercise history.
- `internal/`: Core logic and data management.
    - `exercise/`: Exercise logging and history management.
    - `git/`: Git integration for data persistence.
    - `notes/`: Note creation and retrieval logic.

## Getting Started

### Prerequisites

- Go (1.21 or later)

### Building the Project

To build the MCP server, run:

```bash
make build
```

The binary will be placed in the `bin/` directory.

### Running the Server

You can run the server using:

```bash
make run
```

The server communicates over standard I/O (stdio).

### Running Tests

To run all unit tests:

```bash
make test
```

## Available Tools (via MCP)

- `add_note`: Create a new note with a title and content.
- `list_notes`: List all stored notes.
- `get_note`: Read a specific note by its ID.
- `log_exercise`: Log an exercise activity with duration and notes.
- `get_exercise_history`: Retrieve the history of logged exercises.
- `ping`: Verify connectivity to the MCP server.

## Contributing

Please refer to `agents/AGENT.md` for guidelines on agent behavior and code style.
