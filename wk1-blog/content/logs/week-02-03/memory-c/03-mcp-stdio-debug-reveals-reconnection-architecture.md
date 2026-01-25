---
title: "Mcp Stdio Debug Reveals Reconnection Architecture"
weight: 3
agent: "Claude Code v2.1.7"
duration: "6m 36s"
---

## TL;DR - Session insights

- **The MCP server worked perfectly in isolation but Claude Code couldn't load it** - Manual testing with JSON-RPC initialization handshake succeeded, proving the server itself was functional. The real issue was at the client-side connection layer.
- **"Failed to reconnect to memory" reveals Claude Code's reconnection architecture** - The `/mcp` command output shows Claude Code attempts to maintain persistent connections to MCP servers and has a reconnection mechanism that can fail independently of server health.
- **Debugging MCP requires testing the full protocol handshake, not just the binary** - Simply checking `which todo` or running the script isn't enough; you need to simulate the actual JSON-RPC `initialize` → `notifications/initialized` → `tools/list` sequence.
- **The memory-service has a multi-tier architecture** - The MCP script is just a bash wrapper that forwards requests to a separate HTTP backend at `localhost:8080/rpc`, adding a potential failure point that isn't obvious from the config.

---

## Session Log

* **USER**: Query memory MCP for something

* **AI**: I'll query the memory MCP servers to see what's stored.

* **TOOL-CALL**: mcp__note__list_projects `{}`

* **TOOL-CALL**: mcp__zet__list_zettels `{"limit":10}`

* **TOOL-CALL**: mcp__todo__get_projects `{}`

* **CLI-MESSAGE**: `The user doesn't want to proceed with this tool use. The tool use was rejected (eg. if it was a file edit, the new_string was NOT written to the file). STOP what you are doing and wait for the user to`

* **CLI-MESSAGE**: `The user doesn't want to take this action right now. STOP what you are doing and wait for the user to tell you how to proceed.`

* **CLI-MESSAGE**: `The user doesn't want to take this action right now. STOP what you are doing and wait for the user to tell you how to proceed.`

* **USER**: [Request interrupted by user for tool use]

* **USER**: Looks like memory MCP loading failed.

* **AI**: That explains it. The MCP servers for notes (`mcp__note__`), zettels (`mcp__zet__`), and todos (`mcp__todo__`) aren't loading properly. 
  
  This is typically a configuration issue. You may want to check:
  
  1. **MCP server config** - Verify your `~/.claude/claude_desktop_config.json` or similar config file has the correct server definitions
  2. **Server processes** - Ensure the MCP server processes are running
  3. **Dependencies** - Check if the servers have any missing dependencies
  
  Would you like help troubleshooting the MCP configuration?

* **USER**: yes. Look at ~/.claude.json

* **TOOL-CALL**: Read `{"file_path":"~/.claude.json"}`

* **TOOL**:
  ```
       1→{
       2→  "numStartups": 12,
       3→  "tipsHistory": {
       4→    "new-user-warmup": 8,
       5→    "plan-mode-for-complex-tasks": 8,
       6→    "terminal-setup": 7,
       7→    "memory-command": 7,
       8→    "theme-command": 7,
       9→    "status-line": 8,
      10→    "stickers-command": 8,
      11→    "prompt-queue": 8,
      12→    "enter-to-steer-in-relatime": 8,
      13→    "todo-list": 8,
      14→    "ide-upsell-external-terminal": 12,
      15→    "install-github-app": 8,
      16→    "install-slack-app": 8,
      17→    "drag-and-drop-images": 8,
      18→    "paste-images-mac": 8,
      19→    "double-esc-code-restore": 8,
      20→    "continue": 8,
      21→    "shift-tab": 8,
      22→    "image-paste": 8,
      23→    "ultrathink-keyword": 8,
      24→    "custom-agents": 8,
      25→    "desktop-app": 8,
      26→    "web-app": 8,
      27→    "mobile-app": 8,
      28→    "guest-passes": 11,
      29→    "permissions": 11,
      30→    "rename-conversation": 11,
  [...238 more lines...]
  ```

* **AI**: I can see your MCP server config (lines 233-257). You have four servers configured:
  
  | Server | Command |
  |--------|---------|
  | `todo` | `todo mcp` |
  | `note` | `note mcp` |
  | `zet` | `zet mcp` |
  | `memory` | `./grid/memory-service/scripts/memory-mcp` |
  
  The `todo`, `note`, and `zet` servers expect binaries in your PATH. Let me check if they exist.

* **TOOL-CALL**: Bash `{"command":"which todo note zet 2>&1; echo \"---\"; ls -la ./grid/memory-service/scripts/memory-mcp 2>&1","description":"Check if MCP server binaries exist in PATH"}`

* **TOOL**:
  ```
  ~/go/bin/todo
  ~/go/bin/note
  ~/go/bin/zet
  ---
  -rwxr-xr-x@ 1 user  staff  10708 Jan 13 22:18 ./grid/memory-service/scripts/memory-mcp
  ```

* **AI**: The binaries all exist. Let me test if they run correctly.

* **TOOL-CALL**: Bash `{"command":"echo '{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"initialize\",\"params\":{\"protocolVersion\":\"2024-11-05\",\"capabilities\":{},\"clientInfo\":{\"name\":\"test\",\"version\":\"1.0\"}}}' | timeout 3 ~/go/bin/todo mcp 2>&1 || echo \"Exit code: $?\"","description":"Test todo MCP server with initialize request"}`

* **TOOL**:
  ```
  /opt/homebrew/bin/bash: line 1: timeout: command not found
  Exit code:
  ```

* **TOOL-CALL**: Bash `{"command":"echo '{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"initialize\",\"params\":{\"protocolVersion\":\"2024-11-05\",\"capabilities\":{},\"clientInfo\":{\"name\":\"test\",\"version\":\"1.0\"}}}' | ~/go/bin/todo mcp 2>&1 &\npid=$!\nsleep 2\nkill $pid 2>/dev/null\nwait $pid 2>/dev/null","description":"Test todo MCP server with initialize request","timeout":5000}`

* **CLI-MESSAGE**: `Exit code 127
2026/01/13 22:30:52 server is closing: EOF`

* **AI**: The server ran but closed on EOF (expected). Let me check if there are any startup errors or missing dependencies.

* **TOOL-CALL**: Bash `{"command":"~/go/bin/todo mcp 2>&1 ","description":"Check todo MCP server startup errors","timeout":3000}`

* **TOOL**: ``

* **USER**: [Request interrupted by user]

* **USER**: Don't test todo, zet and note. They are all working. My problem is only with memory-mcp

* **AI**: Let me test the memory MCP server.

* **TOOL-CALL**: Bash `{"command":"./grid/memory-service/scripts/memory-mcp 2>&1 &\npid=$!\nsleep 2\nkill $pid 2>/dev/null\nwait $pid 2>/dev/null","description":"Test memory MCP server startup","timeout":5000}`

* **TOOL-CALL**: Read `{"file_path":"./grid/memory-service/scripts/memory-mcp"}`

* **CLI-MESSAGE**: `Exit code 143
Command timed out after 5s
[memory-mcp] Started (endpoint: http://localhost:8080/rpc)`

* **TOOL**:
  ```
       1→#!/usr/bin/env bash
       2→#
       3→# memory-mcp - MCP server for memory-service
       4→#
       5→# Exposes memory-service operations as MCP tools for AI assistants.
       6→#
       7→# Environment:
       8→#   MEMORY_SERVICE_HOST  Service endpoint (default: localhost:8080)
       9→#
      10→
      11→set -u
      12→
      13→HOST="${MEMORY_SERVICE_HOST:-localhost:8080}"
      14→ENDPOINT="http://${HOST}/rpc"
      15→
      16→# Log to stderr (not visible to MCP client)
      17→log() {
      18→    echo "[memory-mcp] $*" >&2
      19→}
      20→
      21→# Send JSON-RPC response
      22→respond() {
      23→    local id="$1"
      24→    local result="$2"
      25→    printf '{"jsonrpc":"2.0","id":%s,"result":%s}\n' "$id" "$result"
      26→}
      27→
      28→# Send JSON-RPC error
      29→respond_error() {
      30→    local id="$1"
  [...349 more lines...]
  ```

* **AI**: The MCP script itself starts fine. It connects to a backend at `http://localhost:8080/rpc`. Let me check if that service is running.

* **TOOL-CALL**: Bash `{"command":"curl -s -X POST \"http://localhost:8080/rpc\" -H \"Content-Type: application/json\" -d '{\"jsonrpc\":\"2.0\",\"method\":\"list_sessions\",\"params\":{},\"id\":1}' 2>&1","description":"Test if memory-service backend is running"}`

* **TOOL**: `{"jsonrpc":"2.0","id":1,"result":{"sessions":[{"node_id":0,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1},{"node_id":2,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1},{"node_id":4,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1},{"node_id":6,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1}],"total":4}}`

* **AI**: The backend is running and responding correctly. The issue is that I don't see any `mcp__memory__*` tools in my available tools - the MCP server isn't being loaded by Claude Code even though it's configured.
  
  Let me test the full MCP initialization handshake.

* **TOOL-CALL**: Bash `{"command":"echo '{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"initialize\",\"params\":{\"protocolVersion\":\"2024-11-05\",\"capabilities\":{},\"clientInfo\":{\"name\":\"test\",\"version\":\"1.0\"}}}\n{\"jsonrpc\":\"2.0\",\"method\":\"notifications/initialized\"}\n{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"tools/list\",\"params\":{}}' | MEMORY_SERVICE_HOST=localhost:8080 ./grid/memory-service/scripts/memory-mcp 2>&1","description":"Test full MCP handshake with...`

* **TOOL**:
  ```
  [memory-mcp] Started (endpoint: http://localhost:8080/rpc)
  [memory-mcp] Request: {"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}
  {"jsonrpc":"2.0","id":1,"result":{
                  "protocolVersion": "2024-11-05",
                  "capabilities": {
                      "tools": {}
                  },
                  "serverInfo": {
                      "name": "memory-mcp",
                      "version": "1.0.0"
                  }
              }}
  [memory-mcp] Request: {"jsonrpc":"2.0","method":"notifications/initialized"}
  [memory-mcp] Request: {"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}
  {"jsonrpc":"2.0","id":2,"result":{
    "tools": [
      {
        "name": "memory_query",
        "description": "Semantic search across stored memories. Returns relevant content ranked by similarity score.",
        "inputSchema": {
          "type": "object",
          "properties": {
            "query": {
              "type": "string",
              "description": "Search text to find semantically similar content"
            },
            "level": {
              "type": "string",
              "enum": ["session", "message", "block", "statement"],
              "description": "Search only this specific level"
  [...138 more lines...]
  ```

* **AI**: The MCP server works perfectly when tested manually - initialization, handshake, and tools list all return correctly.
  
  Since the script works but Claude Code isn't loading it, the issue is likely:
  
  1. **Claude Code wasn't restarted** after you added/changed the config
  2. **Startup timing** - the server might be timing out during Claude's startup
  
  Try running `/mcp` in Claude Code to see the MCP server status, or restart Claude Code with:
  
  ```bash
  claude --mcp-debug
  ```
  
  This will show detailed MCP connection logs. Alternatively, you can check if there's a Claude Code log file.

* **TOOL-CALL**: Bash `{"command":"ls -la ~/Library/Logs/Claude\\ Code/ 2>/dev/null || ls -la ~/.claude/logs/ 2>/dev/null || echo \"No log directory found\"","description":"Check for Claude Code logs"}`

* **TOOL**: `No log directory found`

* **AI**: No logs found. Can you run `/mcp` to check the MCP server status? That should show if the memory server is loaded or errored.

* **USER**: <local-command-caveat>Caveat: The messages below were generated by the user while running local commands. DO NOT respond to these messages or otherwise consider them in your response unless the user explicitly asks you to.</local-command-caveat>

* **USER**: <command-name>/mcp</command-name>
              <command-message>mcp</command-message>
              <command-args></command-args>

* **USER**: <local-command-stdout>Failed to reconnect to memory.</local-command-stdout>
