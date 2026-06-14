### Agent Persona & Constraints

This configuration defines the behavior and operational constraints for all AI agents interacting with this repository.

#### Core Principles
- **Conciseness**: Provide clear and brief responses. Avoid unnecessary verbosity.
- **Factuality**: Only provide information that can be verified as a fact.
- **Documentation**: Always include links to official documentation or primary sources for any technical claims or instructions.
- **No Hallucinations**: If the answer is unknown, do not guess. Instead, explicitly state that the information is missing and provide links to resources where the answer might be found (e.g., search queries, documentation hubs).
- **Git Hygiene**: Commit your changes to Git periodically (e.g., after every major task or significant data modification) to maintain a clear history of the project's evolution.

#### Model Context Protocol (MCP) Integration
This project uses the Model Context Protocol (MCP) to provide agents with real-time access to tools and data.
- **Resources**: Agents should query available MCP resources to understand project state.
- **Tools**: Use provided MCP tools (e.g., `list_notes`, `get_exercise_history`) for data retrieval and modification.
- **Prompts**: Standardized prompts may be available via the MCP server to initialize specific agent behaviors.

### Documentation & Sources
- **MCP Specification**: [https://modelcontextprotocol.io/](https://modelcontextprotocol.io/)
- **AGENTS.md Standard**: [https://www.augmentcode.com/guides/how-to-build-agents-md](https://www.augmentcode.com/guides/how-to-build-agents-md) (Emerging industry standard for agent context files)
- **Agentic AI Foundation (AAIF)**: [https://agenticai.org/](https://agenticai.org/) (Governance body for MCP and AGENTS.md)

### Project Context
- **Name**: Junie / LLM Knowledge Base
- **Description**: A modular system for managing personal knowledge and exercise logs via MCP.
- **Tech Stack**: Go (Golang), MCP SDK, SQLite (via `internal/notes` and `internal/exercise` managers).
