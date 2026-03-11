package lsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// jsonRPCMessage represents a JSON-RPC 2.0 message.
type jsonRPCMessage struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      *int64           `json:"id,omitempty"`
	Method  string           `json:"method,omitempty"`
	Params  json.RawMessage  `json:"params,omitempty"`
	Result  json.RawMessage  `json:"result,omitempty"`
	Error   *jsonRPCError    `json:"error,omitempty"`
}

// jsonRPCError is a JSON-RPC error object.
type jsonRPCError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (e *jsonRPCError) Error() string {
	return fmt.Sprintf("LSP error %d: %s", e.Code, e.Message)
}

// Client manages a language server subprocess lifecycle.
type Client struct {
	cmd     *exec.Cmd
	stdin   io.WriteCloser
	stdout  *bufio.Reader
	nextID  atomic.Int64
	pending map[int64]chan json.RawMessage
	mu      sync.Mutex
	done    chan struct{}
	timeout time.Duration
}

// Start spawns a language server subprocess.
func Start(command string, args []string, rootDir string) (*Client, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = rootDir
	cmd.Stderr = os.Stderr // let server errors pass through

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start %s: %w", command, err)
	}

	c := &Client{
		cmd:     cmd,
		stdin:   stdin,
		stdout:  bufio.NewReaderSize(stdout, 1024*1024),
		pending: make(map[int64]chan json.RawMessage),
		done:    make(chan struct{}),
		timeout: 30 * time.Second,
	}

	go c.readLoop()
	return c, nil
}

// SetTimeout sets the per-request timeout.
func (c *Client) SetTimeout(d time.Duration) {
	c.timeout = d
}

// readLoop reads messages from the server and dispatches responses.
func (c *Client) readLoop() {
	defer close(c.done)
	for {
		msg, err := c.readMessage()
		if err != nil {
			return // connection closed
		}
		if msg.ID != nil {
			c.mu.Lock()
			ch, ok := c.pending[*msg.ID]
			if ok {
				delete(c.pending, *msg.ID)
			}
			c.mu.Unlock()
			if ok {
				if msg.Error != nil {
					ch <- nil // signal error
				} else {
					ch <- msg.Result
				}
				close(ch)
			}
		}
		// Notifications (no ID) are ignored for now
	}
}

// readMessage reads a single LSP message with Content-Length framing.
func (c *Client) readMessage() (*jsonRPCMessage, error) {
	// Read headers
	contentLength := 0
	for {
		line, err := c.stdout.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // end of headers
		}
		if strings.HasPrefix(line, "Content-Length:") {
			val := strings.TrimSpace(strings.TrimPrefix(line, "Content-Length:"))
			contentLength, _ = strconv.Atoi(val)
		}
	}

	if contentLength == 0 {
		return nil, fmt.Errorf("missing Content-Length header")
	}

	// Read body
	body := make([]byte, contentLength)
	_, err := io.ReadFull(c.stdout, body)
	if err != nil {
		return nil, err
	}

	var msg jsonRPCMessage
	if err := json.Unmarshal(body, &msg); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return &msg, nil
}

// Request sends a JSON-RPC request and waits for the response.
func (c *Client) Request(method string, params interface{}) (json.RawMessage, error) {
	id := c.nextID.Add(1)

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	msg := jsonRPCMessage{
		JSONRPC: "2.0",
		ID:      &id,
		Method:  method,
		Params:  paramsJSON,
	}

	ch := make(chan json.RawMessage, 1)
	c.mu.Lock()
	c.pending[id] = ch
	c.mu.Unlock()

	if err := c.writeMessage(&msg); err != nil {
		c.mu.Lock()
		delete(c.pending, id)
		c.mu.Unlock()
		return nil, err
	}

	select {
	case result := <-ch:
		if result == nil {
			return nil, fmt.Errorf("LSP error for %s", method)
		}
		return result, nil
	case <-time.After(c.timeout):
		c.mu.Lock()
		delete(c.pending, id)
		c.mu.Unlock()
		return nil, fmt.Errorf("timeout waiting for %s response", method)
	case <-c.done:
		return nil, fmt.Errorf("server closed")
	}
}

// Notify sends a JSON-RPC notification (no response expected).
func (c *Client) Notify(method string, params interface{}) error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}

	msg := jsonRPCMessage{
		JSONRPC: "2.0",
		Method:  method,
		Params:  paramsJSON,
	}
	return c.writeMessage(&msg)
}

// writeMessage writes a JSON-RPC message with Content-Length framing.
func (c *Client) writeMessage(msg *jsonRPCMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, err := io.WriteString(c.stdin, header); err != nil {
		return err
	}
	if _, err := c.stdin.Write(body); err != nil {
		return err
	}
	return nil
}

// Initialize sends the LSP initialize request with capabilities.
func (c *Client) Initialize(rootURI string) (*InitializeResult, error) {
	params := InitializeParams{
		ProcessID: os.Getpid(),
		RootURI:   rootURI,
		Capabilities: ClientCapabilities{
			TextDocument: TextDocumentClientCapabilities{
				DocumentSymbol: &DocumentSymbolCapability{
					HierarchicalDocumentSymbolSupport: true,
				},
				References: &ReferencesCapability{
					DynamicRegistration: false,
				},
				CallHierarchy: &CallHierarchyCapability{
					DynamicRegistration: false,
				},
			},
		},
	}

	result, err := c.Request("initialize", params)
	if err != nil {
		return nil, err
	}

	// Send initialized notification
	if err := c.Notify("initialized", struct{}{}); err != nil {
		return nil, err
	}

	var initResult InitializeResult
	if err := json.Unmarshal(result, &initResult); err != nil {
		return nil, err
	}
	return &initResult, nil
}

// DocumentSymbols returns all symbols in a file.
func (c *Client) DocumentSymbols(uri string) ([]DocumentSymbol, error) {
	params := DocumentSymbolParams{
		TextDocument: TextDocumentIdentifier{URI: uri},
	}
	result, err := c.Request("textDocument/documentSymbol", params)
	if err != nil {
		return nil, err
	}
	var symbols []DocumentSymbol
	if err := json.Unmarshal(result, &symbols); err != nil {
		return nil, err
	}
	return symbols, nil
}

// References finds all references to a symbol at position.
func (c *Client) References(uri string, line, col int) ([]Location, error) {
	params := ReferenceParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{URI: uri},
			Position:     Position{Line: line, Character: col},
		},
		Context: ReferenceContext{IncludeDeclaration: false},
	}
	result, err := c.Request("textDocument/references", params)
	if err != nil {
		return nil, err
	}
	var locations []Location
	if err := json.Unmarshal(result, &locations); err != nil {
		return nil, err
	}
	return locations, nil
}

// CallHierarchyPrepare prepares a call hierarchy for the symbol at position.
func (c *Client) CallHierarchyPrepare(uri string, line, col int) ([]CallHierarchyItem, error) {
	params := CallHierarchyPrepareParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{URI: uri},
			Position:     Position{Line: line, Character: col},
		},
	}
	result, err := c.Request("textDocument/prepareCallHierarchy", params)
	if err != nil {
		return nil, err
	}
	var items []CallHierarchyItem
	if err := json.Unmarshal(result, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// CallHierarchyIncoming returns incoming calls to a symbol.
func (c *Client) CallHierarchyIncoming(item CallHierarchyItem) ([]CallHierarchyIncomingCall, error) {
	params := CallHierarchyIncomingCallsParams{Item: item}
	result, err := c.Request("callHierarchy/incomingCalls", params)
	if err != nil {
		return nil, err
	}
	var calls []CallHierarchyIncomingCall
	if err := json.Unmarshal(result, &calls); err != nil {
		return nil, err
	}
	return calls, nil
}

// CallHierarchyOutgoing returns outgoing calls from a symbol.
func (c *Client) CallHierarchyOutgoing(item CallHierarchyItem) ([]CallHierarchyOutgoingCall, error) {
	params := CallHierarchyOutgoingCallsParams{Item: item}
	result, err := c.Request("callHierarchy/outgoingCalls", params)
	if err != nil {
		return nil, err
	}
	var calls []CallHierarchyOutgoingCall
	if err := json.Unmarshal(result, &calls); err != nil {
		return nil, err
	}
	return calls, nil
}

// DidOpen notifies the server that a document was opened.
func (c *Client) DidOpen(uri, languageID, text string) error {
	return c.Notify("textDocument/didOpen", DidOpenTextDocumentParams{
		TextDocument: TextDocumentItem{
			URI:        uri,
			LanguageID: languageID,
			Version:    1,
			Text:       text,
		},
	})
}

// Shutdown gracefully stops the server.
func (c *Client) Shutdown() error {
	_, err := c.Request("shutdown", nil)
	if err != nil {
		// Best effort — kill the process
		c.cmd.Process.Kill()
		return err
	}
	c.Notify("exit", nil)
	c.stdin.Close()
	return c.cmd.Wait()
}
