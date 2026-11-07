# Lesson 12: MCP Tic-Tac-Toe Server

An MCP (Model Context Protocol) server implemented in Go using standard JSON-RPC 2.0 over standard I/O (stdin/stdout).

## Exposed Tools

1. **`get_board`**: Returns the current 3x3 Tic-Tac-Toe board state.
2. **`make_move`**: Places an `O` at the specified `row` and `col` (0-indexed). Validates bounds and checks for already-taken cells.

## Running & Testing Standalone

Run directly with Go:

```bash
go run main.go
```

You can pipe JSON-RPC requests directly into the server to test it:

```bash
printf '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}\n{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}\n{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"make_move","arguments":{"row":1,"col":1}}}\n' | go run main.go
```

## Hooking Up to Antigravity IDE

In Antigravity IDE, MCP servers are managed via the **Agent Panel**:

1. Look at the top-right corner of the **Agent Panel** (sidebar chat) and click the **`⋮` (More actions)** menu (three vertical dots).
2. Click **Manage MCP Servers**.
3. Click **View raw config** (or directly open/create `~/.gemini/config/mcp_config.json`, or `.agents/mcp_config.json` inside your project root) and add:

```json
{
  "mcpServers": {
    "tictactoe": {
      "command": "go",
      "args": ["run", "/Users/zbadgerm/coin-collector/ai-lessons/lesson-12-mcp/main.go"]
    }
  }
}
```

4. Return to **Manage MCP Servers** to verify that `tictactoe` is connected with `get_board` and `make_move`.
5. In chat, prompt the agent:
   > *"Let's play Tic-Tac-Toe! You are 'O' and I am 'X'. Inspect the board using `get_board` and place your move with `make_move`. You go first!"*
