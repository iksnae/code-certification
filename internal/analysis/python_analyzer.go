package analysis

import (
	"context"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
)

// PythonAnalyzer implements Analyzer for Python using tree-sitter.
type PythonAnalyzer struct {
	parser *sitter.Parser
}

// NewPythonAnalyzer creates a new Python analyzer.
func NewPythonAnalyzer() *PythonAnalyzer {
	p := sitter.NewParser()
	p.SetLanguage(python.GetLanguage())
	return &PythonAnalyzer{parser: p}
}

func init() {
	Register("py", func() Analyzer { return NewPythonAnalyzer() })
}

func (a *PythonAnalyzer) Language() string { return "py" }

// Discover finds all symbols in Python source.
func (a *PythonAnalyzer) Discover(path string, src []byte) ([]Symbol, error) {
	tree, err := a.parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, err
	}
	defer tree.Close()

	root := tree.RootNode()
	var symbols []Symbol

	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		symbols = append(symbols, pyDiscoverNode(child, src, "")...)
	}

	return symbols, nil
}

func pyDiscoverNode(node *sitter.Node, src []byte, parent string) []Symbol {
	var symbols []Symbol

	switch node.Type() {
	case "function_definition":
		name := pyNodeName(node, src)
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
				Exported:  !strings.HasPrefix(name, "_") || name == "__init__",
				Parent:    parent,
			})
		}

	case "class_definition":
		name := pyNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolClass,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  !strings.HasPrefix(name, "_"),
			})
			// Discover methods within class body
			if body := pyChildByType(node, "block"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					child := body.Child(j)
					symbols = append(symbols, pyDiscoverNode(child, src, name)...)
				}
			}
		}

	case "decorated_definition":
		// Unwrap decorators
		for j := 0; j < int(node.ChildCount()); j++ {
			child := node.Child(j)
			if child.Type() == "function_definition" || child.Type() == "class_definition" {
				symbols = append(symbols, pyDiscoverNode(child, src, parent)...)
			}
		}
	}

	return symbols
}

// Analyze returns structural metrics for a symbol.
func (a *PythonAnalyzer) Analyze(path string, src []byte, symbol string) (Metrics, error) {
	tree, err := a.parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return Metrics{}, nil
	}
	defer tree.Close()

	root := tree.RootNode()
	unsafeImports := pyDetectUnsafeImports(root, src)

	node := pyFindSymbol(root, src, symbol)
	if node == nil {
		return Metrics{}, nil
	}

	m := pyAnalyzeNode(node, src)
	m.UnsafeImports = unsafeImports
	return m, nil
}

// AnalyzeFile returns file-level metrics.
func (a *PythonAnalyzer) AnalyzeFile(path string, src []byte) (FileMetrics, error) {
	return FileMetrics{}, nil
}

func pyAnalyzeNode(node *sitter.Node, src []byte) Metrics {
	switch node.Type() {
	case "function_definition":
		return pyAnalyzeFunction(node, src)
	case "class_definition":
		return pyAnalyzeClass(node, src)
	case "decorated_definition":
		for j := 0; j < int(node.ChildCount()); j++ {
			child := node.Child(j)
			if child.Type() == "function_definition" || child.Type() == "class_definition" {
				return pyAnalyzeNode(child, src)
			}
		}
	}
	return Metrics{}
}

func pyAnalyzeFunction(node *sitter.Node, src []byte) Metrics {
	name := pyNodeName(node, src)
	m := Metrics{
		IsExported:    !strings.HasPrefix(name, "_") || name == "__init__",
		HasDocComment: pyHasDocstring(node),
		IsConstructor: name == "__init__" || name == "__new__",
	}

	// Parameters (exclude self/cls)
	if params := pyChildByType(node, "parameters"); params != nil {
		for j := 0; j < int(params.ChildCount()); j++ {
			child := params.Child(j)
			switch child.Type() {
			case "identifier":
				pname := child.Content(src)
				if pname != "self" && pname != "cls" {
					m.ParamCount++
				}
			case "default_parameter", "typed_parameter", "typed_default_parameter":
				pname := ""
				if id := pyChildByType(child, "identifier"); id != nil {
					pname = id.Content(src)
				}
				if pname != "self" && pname != "cls" {
					m.ParamCount++
				}
			case "list_splat_pattern", "dictionary_splat_pattern":
				m.ParamCount++
			}
		}
	}

	// Return type annotation
	if pyChildByType(node, "type") != nil {
		m.ReturnCount = 1
	}

	// Body analysis
	if body := pyChildByType(node, "block"); body != nil {
		m.FuncLines = int(body.EndPoint().Row) - int(body.StartPoint().Row)
		if m.FuncLines < 0 {
			m.FuncLines = 0
		}
		m.MaxNestingDepth = pyComputeNesting(body, 0)
		m.LoopNestingDepth = pyComputeLoopNesting(body, 0)
		m.CyclomaticComplexity = pyComputeCyclomatic(body)
		m.CognitiveComplexity = pyComputeCognitive(body, 0)
		m.PanicCalls = pyCountRaises(body)
		m.EmptyCatchBlocks = pyCountEmptyExcept(body)
		m.RecursiveCalls = pyCountRecursive(body, name, src)
	}

	return m
}

func pyAnalyzeClass(node *sitter.Node, src []byte) Metrics {
	name := pyNodeName(node, src)
	m := Metrics{
		IsExported:    !strings.HasPrefix(name, "_"),
		HasDocComment: pyHasDocstring(node),
	}

	if body := pyChildByType(node, "block"); body != nil {
		for j := 0; j < int(body.ChildCount()); j++ {
			child := body.Child(j)
			if child.Type() == "function_definition" {
				m.MethodCount++
			} else if child.Type() == "decorated_definition" {
				for k := 0; k < int(child.ChildCount()); k++ {
					if child.Child(k).Type() == "function_definition" {
						m.MethodCount++
					}
				}
			}
		}
	}

	return m
}

// --- Helpers ---

func pyNodeName(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "identifier" {
			return child.Content(src)
		}
	}
	return ""
}

func pyChildByType(node *sitter.Node, nodeType string) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == nodeType {
			return child
		}
	}
	return nil
}

func pyHasDocstring(node *sitter.Node) bool {
	body := pyChildByType(node, "block")
	if body == nil {
		return false
	}
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		if child.Type() == "expression_statement" {
			if child.ChildCount() > 0 && child.Child(0).Type() == "string" {
				return true
			}
		}
		// Skip comments, but stop at non-docstring statements
		if child.Type() != "comment" && child.Type() != "expression_statement" {
			break
		}
	}
	return false
}

// --- Nesting ---

func pyComputeNesting(node *sitter.Node, depth int) int {
	maxDepth := depth
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "if_statement", "for_statement", "while_statement", "with_statement", "try_statement":
			d := depth + 1
			if d > maxDepth {
				maxDepth = d
			}
			inner := pyComputeNesting(child, d)
			if inner > maxDepth {
				maxDepth = inner
			}
		default:
			inner := pyComputeNesting(child, depth)
			if inner > maxDepth {
				maxDepth = inner
			}
		}
	}
	return maxDepth
}

func pyComputeLoopNesting(node *sitter.Node, depth int) int {
	maxDepth := 0
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "for_statement", "while_statement":
			d := depth + 1
			if d > maxDepth {
				maxDepth = d
			}
			inner := pyComputeLoopNesting(child, d)
			if inner > maxDepth {
				maxDepth = inner
			}
		default:
			inner := pyComputeLoopNesting(child, depth)
			if inner > maxDepth {
				maxDepth = inner
			}
		}
	}
	return maxDepth
}

// --- Complexity ---

func pyComputeCyclomatic(body *sitter.Node) int {
	complexity := 1
	pyWalkAll(body, func(n *sitter.Node) {
		switch n.Type() {
		case "if_statement", "elif_clause":
			complexity++
		case "for_statement", "while_statement":
			complexity++
		case "except_clause":
			complexity++
		case "boolean_operator":
			complexity++
		case "conditional_expression":
			complexity++
		}
	})
	return complexity
}

func pyComputeCognitive(node *sitter.Node, nesting int) int {
	total := 0
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "if_statement":
			total += 1 + nesting
			total += pyComputeCognitive(child, nesting+1)
		case "elif_clause":
			total += 1
			total += pyComputeCognitive(child, nesting+1)
		case "else_clause":
			total += 1
			total += pyComputeCognitive(child, nesting+1)
		case "for_statement", "while_statement":
			total += 1 + nesting
			total += pyComputeCognitive(child, nesting+1)
		case "try_statement":
			total += pyComputeCognitive(child, nesting)
		case "except_clause":
			total += 1 + nesting
			total += pyComputeCognitive(child, nesting+1)
		case "boolean_operator":
			total += 1
			total += pyComputeCognitive(child, nesting)
		default:
			total += pyComputeCognitive(child, nesting)
		}
	}
	return total
}

// --- Counts ---

func pyCountRaises(body *sitter.Node) int {
	count := 0
	pyWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "raise_statement" {
			count++
		}
	})
	return count
}

func pyCountEmptyExcept(body *sitter.Node) int {
	count := 0
	pyWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "except_clause" {
			block := pyChildByType(n, "block")
			if block != nil {
				hasContent := false
				for j := 0; j < int(block.ChildCount()); j++ {
					child := block.Child(j)
					if child.Type() == "pass_statement" || child.Type() == "comment" {
						continue
					}
					if child.Type() == "expression_statement" {
						// Bare ellipsis (...) counts as empty
						if child.ChildCount() == 1 && child.Child(0).Type() == "ellipsis" {
							continue
						}
					}
					hasContent = true
				}
				if !hasContent {
					count++
				}
			}
		}
	})
	return count
}

func pyCountRecursive(body *sitter.Node, funcName string, src []byte) int {
	if funcName == "" {
		return 0
	}
	count := 0
	pyWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "call" {
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

var dangerousPyImports = map[string]bool{
	"subprocess": true,
	"os":         false, // tracked, not flagged unless os.system
	"pickle":     true,
	"shelve":     true,
	"marshal":    true,
	"eval":       true,
	"exec":       true,
	"ctypes":     true,
}

func pyDetectUnsafeImports(root *sitter.Node, src []byte) []string {
	var unsafe []string
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		switch child.Type() {
		case "import_statement":
			pyWalkAll(child, func(n *sitter.Node) {
				if n.Type() == "dotted_name" || n.Type() == "identifier" {
					name := n.Content(src)
					if flagged, exists := dangerousPyImports[name]; exists && flagged {
						unsafe = append(unsafe, name)
					}
				}
			})
		case "import_from_statement":
			if mod := pyChildByType(child, "dotted_name"); mod != nil {
				name := mod.Content(src)
				if flagged, exists := dangerousPyImports[name]; exists && flagged {
					unsafe = append(unsafe, name)
				}
			}
		}
	}
	return unsafe
}

// --- Find symbol ---

func pyFindSymbol(root *sitter.Node, src []byte, name string) *sitter.Node {
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		switch child.Type() {
		case "function_definition":
			if pyNodeName(child, src) == name {
				return child
			}
		case "class_definition":
			if pyNodeName(child, src) == name {
				return child
			}
			// Search methods
			if body := pyChildByType(child, "block"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					member := body.Child(j)
					if member.Type() == "function_definition" && pyNodeName(member, src) == name {
						return member
					}
					if member.Type() == "decorated_definition" {
						for k := 0; k < int(member.ChildCount()); k++ {
							fn := member.Child(k)
							if fn.Type() == "function_definition" && pyNodeName(fn, src) == name {
								return fn
							}
						}
					}
				}
			}
		case "decorated_definition":
			for j := 0; j < int(child.ChildCount()); j++ {
				inner := child.Child(j)
				if inner.Type() == "function_definition" && pyNodeName(inner, src) == name {
					return inner
				}
				if inner.Type() == "class_definition" && pyNodeName(inner, src) == name {
					return inner
				}
			}
		}
	}
	return nil
}

func pyWalkAll(node *sitter.Node, fn func(*sitter.Node)) {
	fn(node)
	for i := 0; i < int(node.ChildCount()); i++ {
		pyWalkAll(node.Child(i), fn)
	}
}
