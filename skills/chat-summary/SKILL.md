---
name: chat-summary
description: Automatically summarize the current session and save it to the knowledge base using MCP.
---

# Chat Summary Skill

Use this skill to create a concise summary of the current session and save it as a note in the `llm-knowledge-base`.

## Trigger Conditions
- When the user asks to "summarize the chat", "take notes of our discussion", or "save a summary".
- At the end of a complex task to document the outcome.

## Instructions
1. **Analyze the Session**: Review the entire conversation history, focusing on:
   - The primary objective and the result achieved.
   - Key decisions made and technical details implemented.
   - Any pending items or next steps identified.
2. **Format the Summary**: Create a Markdown-formatted summary with the following sections:
   - `### Objective`: A one-sentence description of what was planned.
   - `### Outcome`: What was actually accomplished.
   - `### Key Technical Details`: Specific code changes, configurations, or commands used.
   - `### Next Steps`: Any remaining work or follow-ups.
3. **Save to Knowledge Base**:
   - Use the `add_note` tool from the `llm-knowledge-base` MCP server.
   - **Title**: `Session Summary: [Brief Topic] - [Current Date]` (e.g., `Session Summary: MCP Skill Creation - 2026-06-14`).
   - **Content**: The Markdown summary generated in Step 2.

## Verification
- Confirm with the user that the summary has been saved and provide the note ID if returned by the tool.
