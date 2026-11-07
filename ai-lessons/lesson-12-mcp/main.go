package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Persistence file so game state persists across tool calls and restarts
const stateFilePath = "/tmp/tictactoe_board.json"

// The global game board
var board = [3][3]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

func loadBoard() {
	data, err := os.ReadFile(stateFilePath)
	if err == nil {
		_ = json.Unmarshal(data, &board)
	}
}

func saveBoard() {
	data, err := json.Marshal(board)
	if err == nil {
		_ = os.WriteFile(stateFilePath, data, 0644)
	}
}

func resetBoardState() {
	board = [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	saveBoard()
}

// Check if a player has won or if it is a draw
func checkGameStatus() string {
	lines := [][3][2]int{
		// Rows
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		// Columns
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		// Diagonals
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		p := board[line[0][0]][line[0][1]]
		if p != "-" && p == board[line[1][0]][line[1][1]] && p == board[line[2][0]][line[2][1]] {
			return fmt.Sprintf("Winner: %s!", p)
		}
	}

	// Check draw
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == "-" {
				return "Game in progress."
			}
		}
	}
	return "Game is a draw!"
}

// JSON-RPC 2.0 Structures
type RPCRequest struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	loadBoard()
	scanner := bufio.NewScanner(os.Stdin)

	// Infinite loop waiting for MCP commands via standard input
	for scanner.Scan() {
		var req RPCRequest
		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
			sendError(0, -32700, "Parse error")
			continue
		}

		handleRequest(req)
	}
}

func handleRequest(req RPCRequest) {
	loadBoard()

	switch req.Method {

	case "initialize":
		sendResponse(req.ID, map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"serverInfo": map[string]string{
				"name":    "tictactoe-mcp",
				"version": "1.0.0",
			},
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
		})

	case "tools/list":
		sendResponse(req.ID, map[string]interface{}{
			"tools": []map[string]interface{}{
				{
					"name":        "get_board",
					"description": "Look at the current Tic-Tac-Toe board and status.",
					"inputSchema": map[string]interface{}{
						"type":       "object",
						"properties": map[string]interface{}{},
					},
				},
				{
					"name":        "make_move",
					"description": "Place a piece (O or X) on the board.",
					"inputSchema": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"row":    map[string]interface{}{"type": "integer", "description": "Row index (0-2)"},
							"col":    map[string]interface{}{"type": "integer", "description": "Column index (0-2)"},
							"player": map[string]interface{}{"type": "string", "description": "Player piece: 'O' or 'X' (defaults to 'O')"},
						},
						"required": []string{"row", "col"},
					},
				},
				{
					"name":        "reset_board",
					"description": "Reset the Tic-Tac-Toe board to start a fresh game.",
					"inputSchema": map[string]interface{}{
						"type":       "object",
						"properties": map[string]interface{}{},
					},
				},
			},
		})

	case "tools/call":
		var params struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments"`
		}

		if err := json.Unmarshal(req.Params, &params); err != nil {
			sendError(req.ID, -32602, "Invalid params")
			return
		}

		if params.Name == "get_board" {
			status := checkGameStatus()
			resultText := fmt.Sprintf("%s\nStatus: %s", getBoardState(), status)
			sendResponse(req.ID, map[string]interface{}{
				"content": []map[string]interface{}{
					{"type": "text", "text": resultText},
				},
			})
		} else if params.Name == "reset_board" {
			resetBoardState()
			sendResponse(req.ID, map[string]interface{}{
				"content": []map[string]interface{}{
					{"type": "text", "text": "Board reset successfully.\n" + getBoardState()},
				},
			})
		} else if params.Name == "make_move" {
			rowFloat, okRow := params.Arguments["row"].(float64)
			colFloat, okCol := params.Arguments["col"].(float64)

			if !okRow || !okCol {
				sendToolError(req.ID, "Missing or invalid row/col arguments.")
				return
			}

			row, col := int(rowFloat), int(colFloat)

			player := "O"
			if p, okP := params.Arguments["player"].(string); okP && (p == "X" || p == "O") {
				player = p
			}

			// Validation: Bounds check and empty space check
			if row < 0 || row > 2 || col < 0 || col > 2 {
				sendToolError(req.ID, "Move out of bounds. Must be 0, 1, or 2.")
				return
			}
			if board[row][col] != "-" {
				sendToolError(req.ID, fmt.Sprintf("Space at %d, %d is already taken. Choose another.", row, col))
				return
			}

			// Apply the move and persist
			board[row][col] = player
			saveBoard()

			status := checkGameStatus()
			resultText := fmt.Sprintf("Move placed by %s at %d, %d.\nCurrent board:\n%sStatus: %s", player, row, col, getBoardState(), status)
			sendResponse(req.ID, map[string]interface{}{
				"content": []map[string]interface{}{
					{"type": "text", "text": resultText},
				},
			})
		} else {
			sendError(req.ID, -32601, "Tool not found")
		}

	default:
		sendError(req.ID, -32601, "Method not found")
	}
}

// getBoardState formats the 3x3 array into a string grid
func getBoardState() string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("%s | %s | %s\n", board[i][0], board[i][1], board[i][2]))
	}
	return sb.String()
}

func sendResponse(id int, result interface{}) {
	resp := RPCResponse{
		Jsonrpc: "2.0",
		ID:      id,
		Result:  result,
	}
	bytes, _ := json.Marshal(resp)
	fmt.Println(string(bytes))
}

func sendError(id int, code int, message string) {
	resp := RPCResponse{
		Jsonrpc: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    code,
			Message: message,
		},
	}
	bytes, _ := json.Marshal(resp)
	fmt.Println(string(bytes))
}

func sendToolError(id int, message string) {
	sendResponse(id, map[string]interface{}{
		"isError": true,
		"content": []map[string]interface{}{
			{"type": "text", "text": message},
		},
	})
}
