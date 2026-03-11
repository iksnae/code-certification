package analysis

import (
	"context"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/rust"
)

// RustAnalyzer implements Analyzer for Rust using tree-sitter.
type RustAnalyzer struct {
	parser *sitter.Parser
}

// NewRustAnalyzer creates a new Rust analyzer.
func NewRustAnalyzer() *RustAnalyzer {
	p := sitter.NewParser()
	p.SetLanguage(rust.GetLanguage())
	return &RustAnalyzer{parser: p}
}

func init() {
	Register("rs", func() Analyzer { return NewRustAnalyzer() })
}

func (a *RustAnalyzer) Language() string { return "rs" }

// Discover finds all symbols in Rust source.
func (a *RustAnalyzer) Discover(path string, src []byte) ([]Symbol, error) {
	tree, err := a.parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, err
	}
	defer tree.Close()

	root := tree.RootNode()
	var symbols []Symbol

	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		symbols = append(symbols, rsDiscoverNode(child, src, "")...)
	}

	return symbols, nil
}

func rsDiscoverNode(node *sitter.Node, src []byte, parent string) []Symbol {
	var symbols []Symbol

	switch node.Type() {
	case "function_item":
		name := rsNodeName(node, src)
		if name != "" {
			kind := SymbolFunction
			if parent != "" {
				kind = SymbolMethod
			}
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      kind,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  rsIsPub(node, src),
				Parent:    parent,
			})
		}

	case "struct_item":
		name := rsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolClass,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  rsIsPub(node, src),
			})
		}

	case "enum_item":
		name := rsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolClass,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  rsIsPub(node, src),
			})
		}

	case "trait_item":
		name := rsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolInterface,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  rsIsPub(node, src),
			})
			// Discover methods in trait body
			if body := rsChildByType(node, "declaration_list"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					symbols = append(symbols, rsDiscoverNode(body.Child(j), src, name)...)
				}
			}
		}

	case "impl_item":
		// Find the type being implemented
		typeName := rsImplTypeName(node, src)
		if body := rsChildByType(node, "declaration_list"); body != nil {
			for j := 0; j < int(body.ChildCount()); j++ {
				symbols = append(symbols, rsDiscoverNode(body.Child(j), src, typeName)...)
			}
		}
	}

	return symbols
}

// Analyze returns structural metrics for a symbol.
func (a *RustAnalyzer) Analyze(path string, src []byte, symbol string) (Metrics, error) {
	tree, err := a.parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return Metrics{}, nil
	}
	defer tree.Close()

	root := tree.RootNode()
	unsafeImports := rsDetectUnsafeImports(root, src)

	node := rsFindSymbol(root, src, symbol)
	if node == nil {
		return Metrics{}, nil
	}

	m := rsAnalyzeNode(node, src)
	m.UnsafeImports = unsafeImports
	return m, nil
}

// AnalyzeFile returns file-level metrics.
func (a *RustAnalyzer) AnalyzeFile(path string, src []byte) (FileMetrics, error) {
	return FileMetrics{}, nil
}

func rsAnalyzeNode(node *sitter.Node, src []byte) Metrics {
	switch node.Type() {
	case "function_item":
		return rsAnalyzeFunction(node, src)
	case "struct_item", "enum_item":
		return Metrics{
			IsExported:    rsIsPub(node, src),
			HasDocComment: rsHasDocComment(node, src),
		}
	case "trait_item":
		return rsAnalyzeTrait(node, src)
	}
	return Metrics{}
}

func rsAnalyzeFunction(node *sitter.Node, src []byte) Metrics {
	name := rsNodeName(node, src)
	m := Metrics{
		IsExported:    rsIsPub(node, src),
		HasDocComment: rsHasDocComment(node, src),
		IsConstructor: name == "new" || name == "default",
	}

	// Parameters (exclude self)
	if params := rsChildByType(node, "parameters"); params != nil {
		for j := 0; j < int(params.ChildCount()); j++ {
			child := params.Child(j)
			if child.Type() == "parameter" || child.Type() == "self_parameter" {
				if child.Type() == "self_parameter" {
					continue // skip &self, &mut self
				}
				m.ParamCount++
			}
		}
	}

	// Return type
	for j := 0; j < int(node.ChildCount()); j++ {
		child := node.Child(j)
		if child.Type() == "type_identifier" || child.Type() == "generic_type" ||
			child.Type() == "reference_type" || child.Type() == "scoped_type_identifier" {
			m.ReturnCount = 1
			break
		}
		// Check for -> in the function signature
		if child.Type() == "->" {
			m.ReturnCount = 1
			break
		}
	}

	// Body analysis
	if body := rsChildByType(node, "block"); body != nil {
		m.FuncLines = int(body.EndPoint().Row) - int(body.StartPoint().Row) - 1
		if m.FuncLines < 0 {
			m.FuncLines = 0
		}
		m.MaxNestingDepth = rsComputeNesting(body, 0)
		m.LoopNestingDepth = rsComputeLoopNesting(body, 0)
		m.CyclomaticComplexity = rsComputeCyclomatic(body)
		m.CognitiveComplexity = rsComputeCognitive(body, 0)
		m.PanicCalls = rsCountPanics(body, src)
		m.RecursiveCalls = rsCountRecursive(body, name, src)
	}

	return m
}

func rsAnalyzeTrait(node *sitter.Node, src []byte) Metrics {
	m := Metrics{
		IsExported:    rsIsPub(node, src),
		HasDocComment: rsHasDocComment(node, src),
	}
	if body := rsChildByType(node, "declaration_list"); body != nil {
		for j := 0; j < int(body.ChildCount()); j++ {
			if body.Child(j).Type() == "function_item" {
				m.MethodCount++
			}
		}
	}
	return m
}

// --- Helpers ---

func rsNodeName(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "identifier" || child.Type() == "type_identifier" {
			return child.Content(src)
		}
	}
	return ""
}

func rsChildByType(node *sitter.Node, nodeType string) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == nodeType {
			return child
		}
	}
	return nil
}

func rsIsPub(node *sitter.Node, src []byte) bool {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "visibility_modifier" {
			return true
		}
	}
	return false
}

func rsHasDocComment(node *sitter.Node, src []byte) bool {
	if prev := node.PrevSibling(); prev != nil {
		if prev.Type() == "line_comment" {
			content := prev.Content(src)
			return strings.HasPrefix(content, "///") || strings.HasPrefix(content, "//!")
		}
	}
	return false
}

func rsImplTypeName(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "type_identifier" {
			return child.Content(src)
		}
		if child.Type() == "generic_type" {
			return rsNodeName(child, src)
		}
	}
	return ""
}

// --- Nesting ---

func rsComputeNesting(node *sitter.Node, depth int) int {
	maxDepth := depth
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "if_expression", "for_expression", "while_expression", "loop_expression",
			"match_expression":
			d := depth + 1
			if d > maxDepth {
				maxDepth = d
			}
			inner := rsComputeNesting(child, d)
			if inner > maxDepth {
				maxDepth = inner
			}
		default:
			inner := rsComputeNesting(child, depth)
			if inner > maxDepth {
				maxDepth = inner
			}
		}
	}
	return maxDepth
}

func rsComputeLoopNesting(node *sitter.Node, depth int) int {
	maxDepth := 0
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "for_expression", "while_expression", "loop_expression":
			d := depth + 1
			if d > maxDepth {
				maxDepth = d
			}
			inner := rsComputeLoopNesting(child, d)
			if inner > maxDepth {
				maxDepth = inner
			}
		default:
			inner := rsComputeLoopNesting(child, depth)
			if inner > maxDepth {
				maxDepth = inner
			}
		}
	}
	return maxDepth
}

// --- Complexity ---

func rsComputeCyclomatic(body *sitter.Node) int {
	complexity := 1
	rsWalkAll(body, func(n *sitter.Node) {
		switch n.Type() {
		case "if_expression":
			complexity++
		case "for_expression", "while_expression", "loop_expression":
			complexity++
		case "match_arm":
			complexity++
		case "binary_expression":
			// Check for && or ||
			for j := 0; j < int(n.ChildCount()); j++ {
				c := n.Child(j)
				if c.Type() == "&&" || c.Type() == "||" {
					complexity++
				}
			}
		}
	})
	return complexity
}

func rsComputeCognitive(node *sitter.Node, nesting int) int {
	total := 0
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "if_expression":
			total += 1 + nesting
			total += rsComputeCognitive(child, nesting+1)
		case "for_expression", "while_expression", "loop_expression":
			total += 1 + nesting
			total += rsComputeCognitive(child, nesting+1)
		case "match_expression":
			total += 1 + nesting
			total += rsComputeCognitive(child, nesting+1)
		case "binary_expression":
			for j := 0; j < int(child.ChildCount()); j++ {
				c := child.Child(j)
				if c.Type() == "&&" || c.Type() == "||" {
					total += 1
				}
			}
			total += rsComputeCognitive(child, nesting)
		default:
			total += rsComputeCognitive(child, nesting)
		}
	}
	return total
}

// --- Counts ---

func rsCountPanics(body *sitter.Node, src []byte) int {
	count := 0
	rsWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "macro_invocation" {
			name := rsNodeName(n, src)
			if name == "panic" || name == "unreachable" || name == "todo" {
				count++
			}
		}
		// Also count .unwrap() calls
		if n.Type() == "call_expression" {
			rsWalkAll(n, func(inner *sitter.Node) {
				if inner.Type() == "field_identifier" && inner.Content(src) == "unwrap" {
					count++
				}
			})
		}
	})
	return count
}

func rsCountRecursive(body *sitter.Node, funcName string, src []byte) int {
	if funcName == "" {
		return 0
	}
	count := 0
	rsWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "call_expression" {
			if fn := n.Child(0); fn != nil {
				if fn.Type() == "identifier" && fn.Content(src) == funcName {
					count++
				}
			}
		}
	})
	return count
}

// --- Unsafe imports ---

var dangerousRustImports = map[string]bool{
	"std::process":           true,
	"std::os::unix::process": true,
	"libc":                   true,
}

func rsDetectUnsafeImports(root *sitter.Node, src []byte) []string {
	var unsafe []string
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		if child.Type() == "use_declaration" {
			path := rsUsePath(child, src)
			for pattern := range dangerousRustImports {
				if strings.HasPrefix(path, pattern) {
					unsafe = append(unsafe, path)
					break
				}
			}
		}
	}
	return unsafe
}

func rsUsePath(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "scoped_identifier", "use_wildcard", "use_list", "use_as_clause", "scoped_use_list":
			return child.Content(src)
		case "identifier":
			return child.Content(src)
		}
	}
	return ""
}

// --- Find symbol ---

func rsFindSymbol(root *sitter.Node, src []byte, name string) *sitter.Node {
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		switch child.Type() {
		case "function_item":
			if rsNodeName(child, src) == name {
				return child
			}
		case "struct_item", "enum_item":
			if rsNodeName(child, src) == name {
				return child
			}
		case "trait_item":
			if rsNodeName(child, src) == name {
				return child
			}
			if body := rsChildByType(child, "declaration_list"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					fn := body.Child(j)
					if fn.Type() == "function_item" && rsNodeName(fn, src) == name {
						return fn
					}
				}
			}
		case "impl_item":
			if body := rsChildByType(child, "declaration_list"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					fn := body.Child(j)
					if fn.Type() == "function_item" && rsNodeName(fn, src) == name {
						return fn
					}
				}
			}
		}
	}
	return nil
}

func rsWalkAll(node *sitter.Node, fn func(*sitter.Node)) {
	fn(node)
	for i := 0; i < int(node.ChildCount()); i++ {
		rsWalkAll(node.Child(i), fn)
	}
}
