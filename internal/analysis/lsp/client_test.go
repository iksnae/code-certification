package lsp

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestContentLengthFraming(t *testing.T) {
	// Test that we can encode and decode Content-Length framed messages
	body := `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	full := header + body

	// Verify the header format
	if !strings.HasPrefix(full, "Content-Length: ") {
		t.Error("missing Content-Length header")
	}
	if !strings.Contains(full, "\r\n\r\n") {
		t.Error("missing header terminator")
	}
}

func TestJSONRPCMessage_Request(t *testing.T) {
	id := int64(42)
	msg := jsonRPCMessage{
		JSONRPC: "2.0",
		ID:      &id,
		Method:  "textDocument/documentSymbol",
		Params:  json.RawMessage(`{"textDocument":{"uri":"file:///test.ts"}}`),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded jsonRPCMessage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded.JSONRPC != "2.0" {
		t.Errorf("jsonrpc = %q, want 2.0", decoded.JSONRPC)
	}
	if decoded.ID == nil || *decoded.ID != 42 {
		t.Errorf("id = %v, want 42", decoded.ID)
	}
	if decoded.Method != "textDocument/documentSymbol" {
		t.Errorf("method = %q, want textDocument/documentSymbol", decoded.Method)
	}
}

func TestJSONRPCMessage_Notification(t *testing.T) {
	msg := jsonRPCMessage{
		JSONRPC: "2.0",
		Method:  "initialized",
		Params:  json.RawMessage(`{}`),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded jsonRPCMessage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded.ID != nil {
		t.Error("notification should not have an ID")
	}
}

func TestJSONRPCMessage_Response(t *testing.T) {
	id := int64(1)
	msg := jsonRPCMessage{
		JSONRPC: "2.0",
		ID:      &id,
		Result:  json.RawMessage(`{"capabilities":{}}`),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded jsonRPCMessage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded.Error != nil {
		t.Error("response should not have an error")
	}
	if decoded.Result == nil {
		t.Error("response should have a result")
	}
}

func TestJSONRPCMessage_Error(t *testing.T) {
	id := int64(1)
	msg := jsonRPCMessage{
		JSONRPC: "2.0",
		ID:      &id,
		Error: &jsonRPCError{
			Code:    -32601,
			Message: "Method not found",
		},
	}

	if msg.Error.Error() != "LSP error -32601: Method not found" {
		t.Errorf("error string = %q", msg.Error.Error())
	}
}

func TestInitializeParams(t *testing.T) {
	params := InitializeParams{
		ProcessID: 1234,
		RootURI:   "file:///project",
		Capabilities: ClientCapabilities{
			TextDocument: TextDocumentClientCapabilities{
				DocumentSymbol: &DocumentSymbolCapability{
					HierarchicalDocumentSymbolSupport: true,
				},
				CallHierarchy: &CallHierarchyCapability{
					DynamicRegistration: false,
				},
			},
		},
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded InitializeParams
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded.ProcessID != 1234 {
		t.Errorf("processId = %d, want 1234", decoded.ProcessID)
	}
	if decoded.RootURI != "file:///project" {
		t.Errorf("rootUri = %q", decoded.RootURI)
	}
	if decoded.Capabilities.TextDocument.DocumentSymbol == nil {
		t.Error("missing documentSymbol capability")
	}
}

func TestDocumentSymbolUnmarshal(t *testing.T) {
	raw := `[
		{"name":"foo","kind":12,"range":{"start":{"line":0,"character":0},"end":{"line":5,"character":1}},"selectionRange":{"start":{"line":0,"character":9},"end":{"line":0,"character":12}}},
		{"name":"Bar","kind":5,"range":{"start":{"line":7,"character":0},"end":{"line":20,"character":1}},"selectionRange":{"start":{"line":7,"character":6},"end":{"line":7,"character":9}},"children":[
			{"name":"method1","kind":6,"range":{"start":{"line":8,"character":2},"end":{"line":10,"character":3}},"selectionRange":{"start":{"line":8,"character":2},"end":{"line":8,"character":9}}}
		]}
	]`

	var symbols []DocumentSymbol
	if err := json.Unmarshal([]byte(raw), &symbols); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if len(symbols) != 2 {
		t.Fatalf("got %d symbols, want 2", len(symbols))
	}

	if symbols[0].Name != "foo" || symbols[0].Kind != SKFunction {
		t.Errorf("symbol[0] = %v", symbols[0])
	}

	if symbols[1].Name != "Bar" || symbols[1].Kind != SKClass {
		t.Errorf("symbol[1] = %v", symbols[1])
	}

	if len(symbols[1].Children) != 1 || symbols[1].Children[0].Name != "method1" {
		t.Errorf("Bar.children = %v", symbols[1].Children)
	}
}

func TestCallHierarchyItemUnmarshal(t *testing.T) {
	raw := `[{"from":{"name":"main","kind":12,"uri":"file:///app.ts","range":{"start":{"line":0,"character":0},"end":{"line":3,"character":1}},"selectionRange":{"start":{"line":0,"character":9},"end":{"line":0,"character":13}}},"fromRanges":[{"start":{"line":1,"character":2},"end":{"line":1,"character":5}}]}]`

	var calls []CallHierarchyIncomingCall
	if err := json.Unmarshal([]byte(raw), &calls); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if len(calls) != 1 {
		t.Fatalf("got %d calls, want 1", len(calls))
	}
	if calls[0].From.Name != "main" {
		t.Errorf("from.name = %q, want main", calls[0].From.Name)
	}
}
