# Lesson 12: Model Context Protocol (MCP) in Action - Tic-Tac-Toe

**Duration:** 1.5 Hours
**Objective:** Build a complete Model Context Protocol (MCP) server in Go that allows an AI to play Tic-Tac-Toe against you by exposing the game state and move commands as standardized tools, and connect it to Antigravity IDE.
**Recommended Reading:** [Introducing the Model Context Protocol](https://www.anthropic.com/news/model-context-protocol) 

## 1. Concept: Games as Servers (30 mins)
* Up until now, we wrote custom JSON schemas for our AI tools. 
* Introduce MCP: A universal standard (like USB-C) for connecting AI models to local applications.
* **The Architecture:** In this project, the Go program is the *Server* (it holds the rules and the 3x3 board). The AI is the *Client*. Instead of a human typing moves, the AI will use MCP tools to "look" at the board and "place" an X or O.

## 2. Activity: The Game State Scaffold & MCP Server (30 mins)
* Before the AI can play, we need a game. Set up a simple 3x3 string array in Go to represent the board.
* Write a quick helper function to print the board to the terminal so the human can see what's happening.

**Example Go Scaffold (The Game State):**
```go
package main

import (
	"fmt"
	"strings"
)

// The global game board
var board = [3][3]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

// Helper to format the board for the AI
func getBoardState() string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("%s | %s | %s\n", board[i][0], board[i][1], board[i][2]))
	}
	return sb.String()
}
```

* Now wire the scaffold into a standard MCP server using JSON-RPC 2.0 over standard I/O (`os.Stdin` and `os.Stdout`).
* The full runnable MCP server is implemented in [`lesson-12-mcp/main.go`](lesson-12-mcp/main.go):
  * `initialize`: Returns protocol version (`2024-11-05`), server identity, and capabilities.
  * `tools/list`: Registers `get_board` and `make_move`.
  * `tools/call`: Executes board checks and validated moves on the board.

**Verify via Terminal:**
```bash
printf '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}\n{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}\n{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"make_move","arguments":{"row":1,"col":1}}}\n' | go run lesson-12-mcp/main.go
```

## 3. Activity: Hooking Up the LLM in Antigravity IDE (30 mins)
* Antigravity IDE natively supports MCP servers using the Stdio transport.
* Connect our Tic-Tac-Toe server to Antigravity so the AI assistant can see and play the game!

### Step 1: Locate MCP in the Antigravity IDE UI
In Antigravity IDE, MCP servers are managed inside the **Agent Panel** (sidebar chat):
1. Open the **Agent Panel** by clicking the Antigravity icon in the left-hand Activity Bar.
2. At the top-right corner of the Agent Panel header, click the **More Actions menu (`⋮` - three vertical dots)**.
3. Select **Manage MCP Servers**.
4. From here, you can see all active MCP servers, check discovered tools, or click **View raw config** to open the JSON configuration file.

> **Tip:** You can also open the Command Palette (<kbd>Cmd</kbd>+<kbd>Shift</kbd>+<kbd>P</kbd> on macOS / <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>P</kbd> on Windows/Linux) and search for `Antigravity: Manage MCP Servers`.

### Step 2: Configure the MCP Server
You can configure the server at the project level or globally:

* **Option A (Project-specific):** Create a file named `.agents/mcp_config.json` in your repository root.
* **Option B (Global):** Edit `~/.gemini/config/mcp_config.json` (or click **View raw config** in the Manage MCP Servers menu).

Add the `tictactoe` server configuration:

```json
{
  "mcpServers": {
    "tictactoe": {
      "command": "go",
      "args": ["run", "/absolute/path/to/coin-collector/ai-lessons/lesson-12-mcp/main.go"]
    }
  }
}
```
*(Replace `/absolute/path/to/coin-collector` with your workspace's actual path, e.g. `/Users/yourusername/coin-collector`).*

> [!NOTE] Do students need to run `main.go` in a separate terminal?
> **No!** Because MCP uses the **Stdio transport**, Antigravity IDE (the client) automatically launches `go run ...` as a background child process and handles JSON-RPC communication through standard input/output (`stdin`/`stdout`). You do not need to keep a terminal tab running.


### Step 3: Verify Tool Discovery
Return to **Manage MCP Servers** (via `⋮` in the Agent Panel). You should see `tictactoe` listed with a green status indicator and two tools:
* `get_board`
* `make_move`

### Step 4: Play Against the AI
Start a conversation with the AI agent in Antigravity IDE:

**Example Prompt:**
> "Let's play Tic-Tac-Toe! You are 'O' and I am 'X'. Use your MCP tools to check the board with `get_board` and make your move with `make_move`. You go first!"

**What to Observe:**
* Notice how the LLM autonomously triggers `tools/call` for `get_board`, parses the empty grid, and calls `make_move` with its chosen coordinates (e.g. `{"row": 1, "col": 1}`).
* Take your turn by telling the AI which cell you claim for 'X', and have it respond with its counter-move!

## 4. Bonus Challenge: Full Two-Player Loop
Currently, `make_move` places the AI's piece (`O`). Try extending the MCP server with:
1. **Record Human Move**: Add a tool `record_human_move` (or add a `"player": "X" | "O"` parameter to `make_move`) so the AI can record the student's moves on the server's board too.
2. **Win Condition Check**: Write a `checkWinner()` function in Go that evaluates rows, columns, and diagonals, declaring a winner when someone gets 3-in-a-row.
3. **Reset Board**: Expose a `reset_board` tool allowing the LLM to wipe the grid for a rematch without restarting the server.

