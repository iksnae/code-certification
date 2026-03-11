// Package lsp provides a generic JSON-RPC 2.0 client for communicating
// with Language Server Protocol servers via stdin/stdout.
package lsp

// InitializeParams is the LSP initialize request params.
type InitializeParams struct {
	ProcessID  int                `json:"processId"`
	RootURI    string             `json:"rootUri"`
	Capabilities ClientCapabilities `json:"capabilities"`
}

// ClientCapabilities describes the client's capabilities.
type ClientCapabilities struct {
	TextDocument TextDocumentClientCapabilities `json:"textDocument,omitempty"`
}

// TextDocumentClientCapabilities describes text document capabilities.
type TextDocumentClientCapabilities struct {
	DocumentSymbol *DocumentSymbolCapability `json:"documentSymbol,omitempty"`
	References     *ReferencesCapability     `json:"references,omitempty"`
	CallHierarchy  *CallHierarchyCapability  `json:"callHierarchy,omitempty"`
}

// DocumentSymbolCapability describes document symbol capabilities.
type DocumentSymbolCapability struct {
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport"`
}

// ReferencesCapability describes references capabilities.
type ReferencesCapability struct {
	DynamicRegistration bool `json:"dynamicRegistration"`
}

// CallHierarchyCapability describes call hierarchy capabilities.
type CallHierarchyCapability struct {
	DynamicRegistration bool `json:"dynamicRegistration"`
}

// InitializeResult is the LSP initialize response.
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
}

// ServerCapabilities describes the server's capabilities.
type ServerCapabilities struct {
	TextDocumentSync       interface{} `json:"textDocumentSync,omitempty"`
	DocumentSymbolProvider bool        `json:"documentSymbolProvider,omitempty"`
	ReferencesProvider     bool        `json:"referencesProvider,omitempty"`
	CallHierarchyProvider  bool        `json:"callHierarchyProvider,omitempty"`
}

// Position in a text document (0-indexed).
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// Range in a text document.
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// Location is a position in a URI.
type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

// TextDocumentIdentifier identifies a text document.
type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

// TextDocumentPositionParams identifies a position in a text document.
type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

// ReferenceParams extends TextDocumentPositionParams with context.
type ReferenceParams struct {
	TextDocumentPositionParams
	Context ReferenceContext `json:"context"`
}

// ReferenceContext is the context for a references request.
type ReferenceContext struct {
	IncludeDeclaration bool `json:"includeDeclaration"`
}

// DocumentSymbol represents a symbol in a document.
type DocumentSymbol struct {
	Name           string           `json:"name"`
	Kind           int              `json:"kind"`
	Range          Range            `json:"range"`
	SelectionRange Range            `json:"selectionRange"`
	Children       []DocumentSymbol `json:"children,omitempty"`
}

// SymbolKind constants (subset).
const (
	SKFunction    = 12
	SKMethod      = 6
	SKClass       = 5
	SKInterface   = 11
	SKVariable    = 13
	SKConstant    = 14
	SKConstructor = 9
)

// CallHierarchyItem represents a call hierarchy item.
type CallHierarchyItem struct {
	Name           string `json:"name"`
	Kind           int    `json:"kind"`
	URI            string `json:"uri"`
	Range          Range  `json:"range"`
	SelectionRange Range  `json:"selectionRange"`
}

// CallHierarchyPrepareParams is the params for callHierarchy/prepare.
type CallHierarchyPrepareParams struct {
	TextDocumentPositionParams
}

// CallHierarchyIncomingCallsParams is the params for callHierarchy/incomingCalls.
type CallHierarchyIncomingCallsParams struct {
	Item CallHierarchyItem `json:"item"`
}

// CallHierarchyIncomingCall represents an incoming call.
type CallHierarchyIncomingCall struct {
	From       CallHierarchyItem `json:"from"`
	FromRanges []Range           `json:"fromRanges"`
}

// CallHierarchyOutgoingCallsParams is the params for callHierarchy/outgoingCalls.
type CallHierarchyOutgoingCallsParams struct {
	Item CallHierarchyItem `json:"item"`
}

// CallHierarchyOutgoingCall represents an outgoing call.
type CallHierarchyOutgoingCall struct {
	To         CallHierarchyItem `json:"to"`
	FromRanges []Range           `json:"fromRanges"`
}

// Diagnostic represents a LSP diagnostic.
type Diagnostic struct {
	Range    Range  `json:"range"`
	Severity int    `json:"severity,omitempty"`
	Code     string `json:"code,omitempty"`
	Source   string `json:"source,omitempty"`
	Message  string `json:"message"`
}

// DiagnosticSeverity constants.
const (
	DiagError   = 1
	DiagWarning = 2
	DiagInfo    = 3
	DiagHint    = 4
)

// DidOpenTextDocumentParams is for textDocument/didOpen.
type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// TextDocumentItem is an item for textDocument/didOpen.
type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

// DocumentSymbolParams is for textDocument/documentSymbol.
type DocumentSymbolParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
