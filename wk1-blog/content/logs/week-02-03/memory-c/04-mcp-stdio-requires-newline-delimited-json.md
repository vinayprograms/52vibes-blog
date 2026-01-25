---
title: "Mcp Stdio Requires Newline Delimited Json"
weight: 4
agent: "Claude Code v2.1.7"
duration: "8m 48s"
---

## TL;DR - Session insights

- MCP over stdio requires **newline-delimited JSON** - multi-line JSON responses break the protocol causing 30-second timeouts
- Architecture pattern: bash script acts as MCP bridge (stdio) to backend service (HTTP/JSON-RPC) - keeps C service lean while bash handles protocol translation
- Debug logs at `~/.claude/debug/*.txt` reveal MCP connection issues - timeout message "Connection timeout triggered after 30003ms" pinpointed the response format problem
- Fix required compacting JSON responses to single lines: heredocs with pretty-printed JSON must be piped through `jq -c '.'`

---

## Session Log

* **USER**: Ok. So the mcp script in `./scripts/memory-mcp` is not working. Is it setup to use stdio transport or something else. The actual MCP server is already running (its a binary under `./build/bin/memory-service`

* **TOOL-CALL**: Read `{"file_path":"./grid/memory-service/scripts/memory-mcp"}`

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

* **AI**: The script **is using stdio transport** for MCP communication. It reads JSON-RPC requests from stdin (line 358) and writes responses to stdout.
  
  However, it's acting as a **bridge/proxy**:
  1. MCP client → connects via stdio to this bash script
  2. Bash script → forwards calls via HTTP to the memory-service binary at `http://localhost:8080/rpc`
  
  So the architecture is:
  ```
  MCP Client <--stdio--> memory-mcp (bash) <--HTTP--> memory-service binary
  ```
  
  The script expects the memory-service binary to be running as an HTTP server on port 8080 (configurable via `MEMORY_SERVICE_HOST` env var).
  
  **Potential issues:**
  1. Is the memory-service actually listening on `localhost:8080`? You can test with:
     ```
     curl -X POST http://localhost:8080/rpc -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"list_sessions","params":{},"id":1}'
     ```
  
  2. If memory-service uses a different port, set `MEMORY_SERVICE_HOST=localhost:<port>` when running the script
  
  3. If you want the MCP server to be the binary itself (no bash proxy), the binary would need to implement MCP's stdio protocol directly instead of HTTP
  
  What specific error are you seeing when you try to use it?

* **USER**: 1. It is running at localhost:8080. I got response. I don't want the server to be a MCP. The bash script must bridge between JSON-RPC on memory-service and any AI agent that supports MCP.

* **AI**: Good, so the architecture is correct - the bash script bridges MCP (stdio) to your memory-service (JSON-RPC over HTTP).
  
  What error are you seeing when you try to use the script? For example:
  - Is it failing to start?
  - Not responding to MCP requests?
  - Returning errors from tool calls?
  
  You can test the script directly:
  ```bash
  echo '{"jsonrpc":"2.0","method":"initialize","params":{},"id":1}' | ./scripts/memory-mcp
  ```
  
  This should return an initialize response. What happens?

* **USER**: [memory-mcp] Started (endpoint: http://localhost:8080/rpc)
  [memory-mcp] Request: {"jsonrpc":"2.0","method":"initialize","params":{},"id":1}
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
  [memory-mcp] Shutdown
  

* **AI**: That works. So the script itself is functional.
  
  What's the actual failure? How are you connecting it to an MCP client? For example:
  - Adding it to Claude Code's MCP config?
  - Another AI agent?
  
  One thing I notice: the JSON response has embedded newlines in the result object. Some MCP clients expect strictly one JSON object per line. That could cause parsing issues.
  
  If you're configuring it in Claude Code (`~/.claude/claude_desktop_config.json` or similar), what does your config look like and what error do you see?

* **USER**: The config is at ~/.claude.json

* **TOOL-CALL**: Read `{"file_path":"~/.claude.json"}`

* **TOOL**:
  ```
       1→{
       2→  "numStartups": 13,
       3→  "tipsHistory": {
       4→    "new-user-warmup": 8,
       5→    "plan-mode-for-complex-tasks": 13,
       6→    "terminal-setup": 7,
       7→    "memory-command": 7,
       8→    "theme-command": 7,
       9→    "status-line": 8,
      10→    "stickers-command": 8,
      11→    "prompt-queue": 13,
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
  [...256 more lines...]
  ```

* **AI**: The config looks correct:
  
  ```json
  "memory": {
    "command": "./grid/memory-service/scripts/memory-mcp",
    "env": {
      "MEMORY_SERVICE_HOST": "localhost:8080"
    }
  }
  ```
  
  What error do you see? Run `/mcp` in Claude Code to check the server status, or look at the error when trying to use a memory tool.

* **USER**: <command-name>/mcp</command-name>
              <command-message>mcp</command-message>
              <command-args></command-args>

* **USER**: <local-command-stdout>Failed to reconnect to memory.</local-command-stdout>

* **USER**: <local-command-caveat>Caveat: The messages below were generated by the user while running local commands. DO NOT respond to these messages or otherwise consider them in your response unless the user explicitly asks you to.</local-command-caveat>

* **USER**: Getting a "failed to reconnect". I started claude code in debug mode. Can you check the logs.

* **TOOL-CALL**: Glob `{"pattern":"**/*.log","path":"~/.claude"}`

* **CLI-MESSAGE**: `The user doesn't want to proceed with this tool use. The tool use was rejected (eg. if it was a file edit, the new_string was NOT written to the file). STOP what you are doing and wait for the user to`

* **USER**: [Request interrupted by user for tool use]

* **USER**:  ~/.claude/debug/38f5f271-e1af-4176-a0ee-bda7a3466079.txt has the debug logs

* **TOOL-CALL**: Read `{"file_path":"~/.claude/debug/38f5f271-e1af-4176-a0ee-bda7a3466079.txt"}`

* **TOOL**:
  ```
       1→2026-01-14T06:39:24.454Z [DEBUG] Watching for changes in setting files ~/.claude...
       2→2026-01-14T06:39:24.457Z [DEBUG] [init] configureGlobalMTLS starting
       3→2026-01-14T06:39:24.457Z [DEBUG] [init] configureGlobalMTLS complete
       4→2026-01-14T06:39:24.457Z [DEBUG] [init] configureGlobalAgents starting
       5→2026-01-14T06:39:24.457Z [DEBUG] [init] configureGlobalAgents complete
       6→2026-01-14T06:39:24.461Z [DEBUG] [LSP MANAGER] initializeLspServerManager() called
       7→2026-01-14T06:39:24.461Z [DEBUG] [LSP MANAGER] Created manager instance, state=pending
       8→2026-01-14T06:39:24.462Z [DEBUG] [LSP MANAGER] Starting async initialization (generation 1)
       9→2026-01-14T06:39:24.463Z [DEBUG] [STARTUP] Loading MCP configs...
      10→2026-01-14T06:39:24.479Z [DEBUG] Loading plugin clangd-lsp from source: "./plugins/clangd-lsp"
      11→2026-01-14T06:39:24.479Z [DEBUG] Using provided version for clangd-lsp@claude-plugins-official: 1.0.0
      12→2026-01-14T06:39:24.480Z [DEBUG] Plugin clangd-lsp@claude-plugins-official version 1.0.0 already cached at ~/.claude/plugins/cache/claude-plugins-official/clangd-lsp/1.0.0
      13→2026-01-14T06:39:24.480Z [DEBUG] Copied local plugin clangd-lsp to versioned cache: ~/.claude/plugins/cache/claude-plugins-official/clangd-lsp/1.0.0
      14→2026-01-14T06:39:24.480Z [DEBUG] Plugin clangd-lsp has no entry.skills defined
      15→2026-01-14T06:39:24.480Z [DEBUG] Found 1 plugins (1 enabled, 0 disabled)
      16→2026-01-14T06:39:24.480Z [DEBUG] [STARTUP] MCP configs loaded in 16ms
      17→2026-01-14T06:39:24.481Z [DEBUG] [STARTUP] Running setup()...
      18→2026-01-14T06:39:24.485Z [DEBUG] Stats cache does not exist, returning empty cache
      19→2026-01-14T06:39:24.485Z [DEBUG] Stats cache empty, processing all historical data in background
      20→2026-01-14T06:39:24.488Z [DEBUG] Loading skills from: managed=/Library/Application Support/ClaudeCode/.claude/skills, user=~/.claude/skills, project=[]
      21→2026-01-14T06:39:24.488Z [DEBUG] Passes: Cache stale, returning cached data and refreshing in background
      22→2026-01-14T06:39:24.492Z [DEBUG] Loaded 1 installed plugins from ~/.claude/plugins/installed_plugins.json
      23→2026-01-14T06:39:24.505Z [DEBUG] Error log sink initialized
      24→2026-01-14T06:39:24.506Z [DEBUG] Writing to temp file: ~/.claude/todos/38f5f271-e1af-4176-a0ee-bda7a3466079-agent-38f5f271-e1af-4176-a0ee-bda7a3466079.json.tmp.26265.1768372764506
      25→2026-01-14T06:39:24.506Z [DEBUG] Temp file written successfully, size: 2 bytes
      26→2026-01-14T06:39:24.506Z [DEBUG] Renaming ~/.claude/todos/38f5f271-e1af-4176-a0ee-bda7a3466079-agent-38f5f271-e1af-4176-a0ee-bda7a3466079.json.tmp.26265.1768372764506 to ~/.claude/todos/38f5f271-e1af-4176-a0ee-bda7a3466079-agent-38f5f271-e1af-4176-a0ee-bda7a3466079.json
      27→2026-01-14T06:39:24.506Z [DEBUG] File ~/.claude/todos/38f5f271-e1af-4176-a0ee-bda7a3466079-agent-38f5f271-e1af-4176-a0ee-bda7a3466079.json written atomically
      28→2026-01-14T06:39:24.510Z [DEBUG] getPluginSkills: Processing 1 enabled plugins
      29→2026-01-14T06:39:24.510Z [DEBUG] Checking plugin clangd-lsp: skillsPath=none, skillsPaths=0 paths
      30→2026-01-14T06:39:24.510Z [DEBUG] Total plugin skills loaded: 0
  [...430 more lines...]
  ```

* **AI**: Found the issue. Looking at lines 154-164:
  
  ```
  [memory-mcp] Request: {"method":"initialize",...}
  ```
  Then 30 seconds later:
  ```
  Connection timeout triggered after 30003ms
  ```
  
  The script receives the request but **Claude Code never sees the response**. The problem is the response has embedded newlines:
  
  ```json
  {"jsonrpc":"2.0","id":1,"result":{
                  "protocolVersion": "2024-11-05",
                  ...
              }}
  ```
  
  MCP over stdio expects **newline-delimited JSON** - each message must be a single line. The multi-line heredoc in the `respond` call breaks the protocol.
  
  The fix is to output compact JSON. In the `respond` function or the initialize handler, the result needs to be on a single line:

* **TOOL-CALL**: Edit `{"replace_all":false,"file_path":"./grid/memory-service/scripts/memory-mcp","old_string":"        initialize)\n            respond \"$id\" '{\n                \"protocolVersion\": \"2024-11-05\",\n                \"capabilities\": {\n                    \"tools\": {}\n                },\n                \"serverInfo\": {\n                    \"name\": \"memory-mcp\",\n                    \"version\": \"1.0.0\"\n                }\n            }'\n            ...`

* **TOOL**:
  ```
  The file ./grid/memory-service/scripts/memory-mcp has been updated. Here's the result of running `cat -n` on a snippet of the edited file:
     307→    local method=$(echo "$request" | jq -r '.method')
     308→    local params=$(echo "$request" | jq -c '.params // {}')
     309→
     310→    case "$method" in
     311→        initialize)
     312→            respond "$id" '{"protocolVersion":"2024-11-05","capabilities":{"tools":{}},"serverInfo":{"name":"memory-mcp","version":"1.0.0"}}'
     313→            ;;
     314→
     315→        notifications/initialized)
     316→            # No response needed for notifications
     317→            ;;
  ```

* **AI**: Now the same issue exists in `get_tools()` - it returns multi-line JSON. Let me fix that too:

* **TOOL-CALL**: Read `{"file_path":"./grid/memory-service/scripts/memory-mcp","offset":63,"limit":160}`

* **TOOL**:
  ```
      63→# Tool definitions
      64→get_tools() {
      65→    cat <<'EOF'
      66→{
      67→  "tools": [
      68→    {
      69→      "name": "memory_query",
      70→      "description": "Semantic search across stored memories. Returns relevant content ranked by similarity score.",
      71→      "inputSchema": {
      72→        "type": "object",
      73→        "properties": {
      74→          "query": {
      75→            "type": "string",
      76→            "description": "Search text to find semantically similar content"
      77→          },
      78→          "level": {
      79→            "type": "string",
      80→            "enum": ["session", "message", "block", "statement"],
      81→            "description": "Search only this specific level"
      82→          },
      83→          "top_level": {
      84→            "type": "string",
      85→            "enum": ["session", "message", "block", "statement"],
      86→            "description": "Highest level to search (default: session)"
      87→          },
      88→          "bottom_level": {
      89→            "type": "string",
      90→            "enum": ["session", "message", "block", "statement"],
      91→            "description": "Lowest level to search (default: statement)"
      92→          },
  [...135 more lines...]
  ```

* **AI**: The `get_tools` function returns multi-line JSON. The easiest fix is to pipe it through `jq -c` where it's used. Let me update the `tools/list` handler:

* **TOOL-CALL**: Edit `{"replace_all":false,"file_path":"./grid/memory-service/scripts/memory-mcp","old_string":"        tools/list)\n            respond \"$id\" \"$(get_tools)\"\n            ;;","new_string":"        tools/list)\n            respond \"$id\" \"$(get_tools | jq -c '.')\"\n            ;;"}`

* **TOOL**:
  ```
  The file ./grid/memory-service/scripts/memory-mcp has been updated. Here's the result of running `cat -n` on a snippet of the edited file:
     315→        notifications/initialized)
     316→            # No response needed for notifications
     317→            ;;
     318→
     319→        tools/list)
     320→            respond "$id" "$(get_tools | jq -c '.')"
     321→            ;;
     322→
     323→        tools/call)
     324→            local tool_name=$(echo "$params" | jq -r '.name')
     325→            local tool_args=$(echo "$params" | jq -c '.arguments // {}')
  ```

* **AI**: Now let me check the `tools/call` response - that also needs to be single-line:

* **TOOL-CALL**: Read `{"file_path":"./grid/memory-service/scripts/memory-mcp","offset":323,"limit":30}`

* **TOOL**:
  ```
     323→        tools/call)
     324→            local tool_name=$(echo "$params" | jq -r '.name')
     325→            local tool_args=$(echo "$params" | jq -c '.arguments // {}')
     326→
     327→            local result
     328→            result=$(execute_tool "$tool_name" "$tool_args")
     329→
     330→            local error_check=$(echo "$result" | jq -r '.error // empty')
     331→            if [[ -n "$error_check" ]]; then
     332→                respond "$id" "{\"content\": [{\"type\": \"text\", \"text\": \"Error: $error_check\"}], \"isError\": true}"
     333→            else
     334→                local text=$(echo "$result" | jq -c '.')
     335→                respond "$id" "{\"content\": [{\"type\": \"text\", \"text\": $(echo "$text" | jq -Rs '.')}]}"
     336→            fi
     337→            ;;
     338→
     339→        *)
     340→            respond_error "$id" -32601 "Method not found: $method"
     341→            ;;
     342→    esac
     343→}
     344→
     345→# Main loop - read JSON-RPC requests from stdin
     346→main() {
     347→    log "Started (endpoint: $ENDPOINT)"
     348→
     349→    while IFS= read -r line; do
     350→        # Skip empty lines
     351→        [[ -z "$line" ]] && continue
     352→
  [...5 more lines...]
  ```

* **AI**: The `tools/call` handler already uses `jq -c` to compact the JSON, so that should be fine.
  
  Test it now:
  
  ```bash
  echo '{"jsonrpc":"2.0","method":"initialize","params":{},"id":1}' | ./scripts/memory-mcp
  ```
  
  The response should now be on a single line. Then restart Claude Code and run `/mcp` to check if the memory server connects.

* **USER**: <local-command-caveat>Caveat: The messages below were generated by the user while running local commands. DO NOT respond to these messages or otherwise consider them in your response unless the user explicitly asks you to.</local-command-caveat>

* **USER**: <command-name>/mcp</command-name>
              <command-message>mcp</command-message>
              <command-args></command-args>

* **USER**: <local-command-stdout>Reconnected to memory.</local-command-stdout>

* **USER**: That works. Now run some queries on this memory MCP

* **AI**: Let me try the memory MCP tools.

* **TOOL-CALL**: mcp__memory__memory_list_sessions `{}`

* **TOOL**:
  ```
  {"sessions":[{"node_id":0,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1},{"node_id":2,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1},{"node_id":4,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1},{"node_id":6,"session_id":"claude-session-mpnet","agent_id":"claude-opus-4.5","message_count":1}],"total":4}
  
  ```

* **AI**: There are 4 sessions. Let me try a semantic search:

* **TOOL-CALL**: mcp__memory__memory_query `{"query":"code implementation","max_results":5}`

* **TOOL**:
  ```
  {"results":[{"node_id":1,"level":"message","score":0.5515109300613403,"content":"User request: Look at memory-service documentation in ../grid/memory-service/ directory and write sample messages to the memory service running on localhost:8080","children_count":0},{"node_id":3,"level":"message","score":0.5386258363723755,"content":"Explored memory-service directory structure and read README.md documentation. The service uses JSON-RPC 2.0 over HTTP with hierarchical storage: session -> message -> block -> statement. Key methods: store, store_block, store_statement, query, drill_down, zoom_out, list_sessions, get_session.","children_count":1},{"node_id":8,"level":"block","score":0.526558518409729,"content":"Memory-service API methods discovered: store (message), store_block, store_statement, query (semantic search with level filtering), drill_down (get children), zoom_out (get ancestors), list_sessions, get_session","children_count":1},{"node_id":10,"level":"statement","score":0.5261462926864624,"content":"Search ranking formula: score = 0.6 * relevance + 0.3 * recency + 0.1 * level_boost","children_count":0},{"node_id":9,"level":"block","score":0.5196473002433777,"content":"Hierarchy levels: session (top) -> message -> block -> statement (bottom). Each level independently indexed for semantic search.","children_count":0}],"total_matches":5,"top_level":"session","bottom_level":"statement","truncated":false}
  
  ```

* **AI**: Semantic search works. Let me test hierarchy navigation - drilling down into a session:

* **TOOL-CALL**: mcp__memory__memory_drill_down `{"id":3}`

* **TOOL**:
  ```
  {"node_id":3,"level":"message","children":[{"node_id":8,"level":"block","content":"Memory-service API methods discovered: store (message), store_block, store_statement, query (semantic search with level filtering), drill_down (get children), zoom_out (get ancestors), list_sessions, get_session","children_count":1},{"node_id":9,"level":"block","content":"Hierarchy levels: session (top) -> message -> block -> statement (bottom). Each level independently indexed for semantic search.","children_count":0}],"children_count":2,"total_children":2}
  
  ```

* **TOOL-CALL**: mcp__memory__memory_zoom_out `{"id":10}`

* **TOOL**:
  ```
  {"node_id":10,"level":"statement","content":"Search ranking formula: score = 0.6 * relevance + 0.3 * recency + 0.1 * level_boost","ancestors":[{"node_id":8,"level":"block","content":"Memory-service API methods discovered: store (message), store_block, store_statement, query (semantic search with level filtering), drill_down (get children), zoom_out (get ancestors), list_sessions, get_session"},{"node_id":3,"level":"message","content":"Explored memory-service directory structure and read README.md documentation. The service uses JSON-RPC 2.0 over HTTP with hierarchical storage: session -> message -> block -> statement. Key methods: store, store_block, store_statement, query, drill_down, zoom_out, list_sessions, get_session."},{"node_id":2,"level":"session"}],"siblings":[]}
  
  ```

* **AI**: All working:
  
  | Tool | Result |
  |------|--------|
  | `list_sessions` | Found 4 sessions |
  | `query` | Semantic search returned 5 ranked results |
  | `drill_down` | Navigated from message → 2 child blocks |
  | `zoom_out` | Traced statement 10 up through block → message → session |
  
  The MCP bridge is functioning correctly. The fix was ensuring all JSON-RPC responses are output on a single line (newline-delimited JSON protocol).
