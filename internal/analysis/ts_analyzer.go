package analysis

import (
	"context"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

// TSAnalyzer implements Analyzer for TypeScript using tree-sitter.
type TSAnalyzer struct {
	parser *sitter.Parser
}

// NewTSAnalyzer creates a new TypeScript analyzer.
func NewTSAnalyzer() *TSAnalyzer {
	p := sitter.NewParser()
	p.SetLanguage(typescript.GetLanguage())
	return &TSAnalyzer{parser: p}
}

func init() {
	Register("ts", func() Analyzer { return NewTSAnalyzer() })
	Register("js", func() Analyzer { return NewTSAnalyzer() }) // JS uses same grammar subset
}

func (a *TSAnalyzer) Language() string { return "ts" }

// Discover finds all symbols in a TypeScript source file.
func (a *TSAnalyzer) Discover(path string, src []byte) ([]Symbol, error) {
	tree, err := a.parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, err
	}
	defer tree.Close()

	root := tree.RootNode()
	var symbols []Symbol

	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		symbols = append(symbols, tsDiscoverNode(child, src, "")...)
	}

	return symbols, nil
}

func tsDiscoverNode(node *sitter.Node, src []byte, parent string) []Symbol {
	var symbols []Symbol

	switch node.Type() {
	case "function_declaration":
		name := tsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolFunction,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  tsIsExported(node, src),
				Parent:    parent,
			})
		}

	case "class_declaration":
		name := tsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolClass,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  tsIsExported(node, src),
			})
			// Discover methods within class body
			if body := tsChildByType(node, "class_body"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					member := body.Child(j)
					if member.Type() == "method_definition" {
						mname := tsNodeName(member, src)
						if mname != "" {
							symbols = append(symbols, Symbol{
								Name:      mname,
								Kind:      SymbolMethod,
								StartLine: int(member.StartPoint().Row) + 1,
								EndLine:   int(member.EndPoint().Row) + 1,
								Exported:  true, // class methods are accessible if class is
								Parent:    name,
							})
						}
					}
				}
			}
		}

	case "interface_declaration":
		name := tsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolInterface,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  tsIsExported(node, src),
			})
		}

	case "type_alias_declaration":
		name := tsNodeName(node, src)
		if name != "" {
			symbols = append(symbols, Symbol{
				Name:      name,
				Kind:      SymbolClass,
				StartLine: int(node.StartPoint().Row) + 1,
				EndLine:   int(node.EndPoint().Row) + 1,
				Exported:  tsIsExported(node, src),
			})
		}

	case "lexical_declaration", "variable_declaration":
		// export const FOO = ...
		for j := 0; j < int(node.ChildCount()); j++ {
			decl := node.Child(j)
			if decl.Type() == "variable_declarator" {
				name := tsNodeName(decl, src)
				if name != "" {
					symbols = append(symbols, Symbol{
						Name:      name,
						Kind:      SymbolConstant,
						StartLine: int(node.StartPoint().Row) + 1,
						EndLine:   int(node.EndPoint().Row) + 1,
						Exported:  tsIsExported(node, src),
					})
				}
			}
		}

	case "export_statement":
		// Unwrap export wrapper
		for j := 0; j < int(node.ChildCount()); j++ {
			child := node.Child(j)
			if child.Type() != "export" && child.Type() != "default" {
				inner := tsDiscoverNode(child, src, parent)
				for k := range inner {
					inner[k].Exported = true
				}
				symbols = append(symbols, inner...)
			}
		}
	}

	return symbols
}

// Analyze returns structural metrics for a specific symbol.
func (a *TSAnalyzer) Analyze(path string, src []byte, symbol string) (Metrics, error) {
	tree, err := a.parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return Metrics{}, nil
	}
	defer tree.Close()

	root := tree.RootNode()

	// Collect file-level context
	unsafeImports := tsDetectUnsafeImports(root, src)

	// Find the symbol node
	node := tsFindSymbol(root, src, symbol)
	if node == nil {
		return Metrics{}, nil
	}

	m := tsAnalyzeNode(node, src)
	m.UnsafeImports = unsafeImports
	return m, nil
}

// AnalyzeFile returns file-level metrics.
func (a *TSAnalyzer) AnalyzeFile(path string, src []byte) (FileMetrics, error) {
	return FileMetrics{}, nil // TS has no init() equivalent
}

// tsAnalyzeNode extracts metrics from a tree-sitter node.
func tsAnalyzeNode(node *sitter.Node, src []byte) Metrics {
	m := Metrics{}

	switch node.Type() {
	case "function_declaration", "arrow_function", "function":
		m = tsAnalyzeFunction(node, src)
	case "method_definition":
		m = tsAnalyzeFunction(node, src)
	case "class_declaration":
		m = tsAnalyzeClass(node, src)
	case "interface_declaration":
		m.IsExported = tsIsExported(node, src)
		m.HasDocComment = tsHasDocComment(node, src)
	case "type_alias_declaration":
		m.IsExported = tsIsExported(node, src)
		m.HasDocComment = tsHasDocComment(node, src)
	case "lexical_declaration", "variable_declaration", "variable_declarator":
		m.IsExported = tsIsExported(node, src)
	case "export_statement":
		// Unwrap and analyze inner declaration
		for i := 0; i < int(node.ChildCount()); i++ {
			child := node.Child(i)
			if child.Type() != "export" && child.Type() != "default" {
				m = tsAnalyzeNode(child, src)
				m.IsExported = true
				return m
			}
		}
	}

	return m
}

func tsAnalyzeFunction(node *sitter.Node, src []byte) Metrics {
	m := Metrics{
		IsExported:    tsIsExported(node, src),
		HasDocComment: tsHasDocComment(node, src),
	}

	name := tsNodeName(node, src)
	m.IsConstructor = name == "constructor" || strings.HasPrefix(name, "create") || strings.HasPrefix(name, "new")

	// Parameters
	if params := tsChildByType(node, "formal_parameters"); params != nil {
		m.ParamCount = tsCountParams(params)
	}

	// Return type annotation
	if tsChildByType(node, "type_annotation") != nil {
		m.ReturnCount = 1
	} else {
		// Check for return statements
		if body := tsFuncBody(node); body != nil {
			if tsHasReturnWithValue(body) {
				m.ReturnCount = 1
			}
		}
	}

	// Body analysis
	if body := tsFuncBody(node); body != nil {
		m.FuncLines = int(body.EndPoint().Row) - int(body.StartPoint().Row) - 1
		if m.FuncLines < 0 {
			m.FuncLines = 0
		}
		m.MaxNestingDepth = tsComputeNesting(body, 0)
		m.LoopNestingDepth = tsComputeLoopNesting(body, 0)
		m.CyclomaticComplexity = tsComputeCyclomatic(body)
		m.CognitiveComplexity = tsComputeCognitive(body, 0)
		m.PanicCalls = tsCountThrows(body)
		m.EmptyCatchBlocks = tsCountEmptyCatch(body)
		m.RecursiveCalls = tsCountRecursive(body, name, src)
	}

	return m
}

func tsAnalyzeClass(node *sitter.Node, src []byte) Metrics {
	m := Metrics{
		IsExported:    tsIsExported(node, src),
		HasDocComment: tsHasDocComment(node, src),
	}

	if body := tsChildByType(node, "class_body"); body != nil {
		for i := 0; i < int(body.ChildCount()); i++ {
			child := body.Child(i)
			if child.Type() == "method_definition" {
				m.MethodCount++
			}
		}
	}

	return m
}

// --- Helper functions ---

func tsNodeName(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "identifier" || child.Type() == "property_identifier" || child.Type() == "type_identifier" {
			return child.Content(src)
		}
	}
	return ""
}

func tsChildByType(node *sitter.Node, nodeType string) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == nodeType {
			return child
		}
	}
	return nil
}

func tsFuncBody(node *sitter.Node) *sitter.Node {
	return tsChildByType(node, "statement_block")
}

func tsIsExported(node *sitter.Node, src []byte) bool {
	// Check if parent is export_statement
	if parent := node.Parent(); parent != nil {
		if parent.Type() == "export_statement" {
			return true
		}
	}
	return false
}

func tsHasDocComment(node *sitter.Node, src []byte) bool {
	// Look for comment node immediately preceding this node
	if node.PrevSibling() != nil {
		prev := node.PrevSibling()
		if prev.Type() == "comment" {
			content := prev.Content(src)
			return strings.HasPrefix(content, "/**") || strings.HasPrefix(content, "///")
		}
	}
	// Check parent (export_statement wraps the declaration)
	if parent := node.Parent(); parent != nil && parent.Type() == "export_statement" {
		if parent.PrevSibling() != nil {
			prev := parent.PrevSibling()
			if prev.Type() == "comment" {
				content := prev.Content(src)
				return strings.HasPrefix(content, "/**") || strings.HasPrefix(content, "///")
			}
		}
	}
	return false
}

func tsCountParams(params *sitter.Node) int {
	count := 0
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		switch child.Type() {
		case "required_parameter", "optional_parameter", "rest_parameter":
			count++
		case "identifier":
			count++
		}
	}
	return count
}

func tsHasReturnWithValue(body *sitter.Node) bool {
	found := false
	tsWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "return_statement" && n.ChildCount() > 1 {
			found = true
		}
	})
	return found
}

// --- Nesting ---

func tsComputeNesting(node *sitter.Node, depth int) int {
	maxDepth := depth
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "if_statement", "for_statement", "for_in_statement", "while_statement",
			"do_statement", "switch_statement", "try_statement":
			d := depth + 1
			if d > maxDepth {
				maxDepth = d
			}
			inner := tsComputeNesting(child, d)
			if inner > maxDepth {
				maxDepth = inner
			}
		default:
			inner := tsComputeNesting(child, depth)
			if inner > maxDepth {
				maxDepth = inner
			}
		}
	}
	return maxDepth
}

func tsComputeLoopNesting(node *sitter.Node, depth int) int {
	maxDepth := 0
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "for_statement", "for_in_statement", "while_statement", "do_statement":
			d := depth + 1
			if d > maxDepth {
				maxDepth = d
			}
			inner := tsComputeLoopNesting(child, d)
			if inner > maxDepth {
				maxDepth = inner
			}
		default:
			inner := tsComputeLoopNesting(child, depth)
			if inner > maxDepth {
				maxDepth = inner
			}
		}
	}
	return maxDepth
}

// --- Complexity ---

func tsComputeCyclomatic(body *sitter.Node) int {
	complexity := 1
	tsWalkAll(body, func(n *sitter.Node) {
		switch n.Type() {
		case "if_statement":
			complexity++
		case "for_statement", "for_in_statement", "while_statement", "do_statement":
			complexity++
		case "switch_case":
			// Only non-default cases
			if tsChildByType(n, "default") == nil {
				complexity++
			}
		case "catch_clause":
			complexity++
		case "ternary_expression":
			complexity++
		case "binary_expression":
			op := ""
			for j := 0; j < int(n.ChildCount()); j++ {
				c := n.Child(j)
				if c.Type() == "&&" || c.Type() == "||" {
					op = c.Type()
				}
			}
			if op == "&&" || op == "||" {
				complexity++
			}
		}
	})
	return complexity
}

func tsComputeCognitive(node *sitter.Node, nesting int) int {
	total := 0
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "if_statement":
			total += 1 + nesting
			// Check for else
			for j := 0; j < int(child.ChildCount()); j++ {
				c := child.Child(j)
				if c.Type() == "else_clause" {
					total += 1
					total += tsComputeCognitive(c, nesting+1)
				} else if c.Type() == "statement_block" || c.Type() == "if_statement" {
					total += tsComputeCognitive(c, nesting+1)
				}
			}

		case "for_statement", "for_in_statement", "while_statement", "do_statement":
			total += 1 + nesting
			total += tsComputeCognitive(child, nesting+1)

		case "switch_statement":
			total += 1 + nesting
			total += tsComputeCognitive(child, nesting+1)

		case "catch_clause":
			total += 1 + nesting
			total += tsComputeCognitive(child, nesting+1)

		case "binary_expression":
			for j := 0; j < int(child.ChildCount()); j++ {
				c := child.Child(j)
				if c.Type() == "&&" || c.Type() == "||" {
					total += 1
				}
			}
			total += tsComputeCognitive(child, nesting)

		default:
			total += tsComputeCognitive(child, nesting)
		}
	}
	return total
}

// --- Specific counts ---

func tsCountThrows(body *sitter.Node) int {
	count := 0
	tsWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "throw_statement" {
			count++
		}
	})
	return count
}

func tsCountEmptyCatch(body *sitter.Node) int {
	count := 0
	tsWalkAll(body, func(n *sitter.Node) {
		if n.Type() == "catch_clause" {
			if block := tsChildByType(n, "statement_block"); block != nil {
				// Count non-trivial children (skip { and })
				stmts := 0
				for j := 0; j < int(block.ChildCount()); j++ {
					c := block.Child(j)
					if c.Type() != "{" && c.Type() != "}" {
						stmts++
					}
				}
				if stmts == 0 {
					count++
				}
			}
		}
	})
	return count
}

func tsCountRecursive(body *sitter.Node, funcName string, src []byte) int {
	if funcName == "" {
		return 0
	}
	count := 0
	tsWalkAll(body, func(n *sitter.Node) {
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

var dangerousTSImports = map[string]bool{
	"child_process": true,
	"eval":          true,
	"vm":            true,
	"cluster":       true,
}

func tsDetectUnsafeImports(root *sitter.Node, src []byte) []string {
	var unsafe []string
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		if child.Type() == "import_statement" {
			// Find the string node (module path)
			tsWalkAll(child, func(n *sitter.Node) {
				if n.Type() == "string" || n.Type() == "string_fragment" {
					val := strings.Trim(n.Content(src), `"'`)
					if dangerousTSImports[val] {
						unsafe = append(unsafe, val)
					}
				}
			})
		}
	}
	return unsafe
}

// --- Tree walker ---

func tsWalkAll(node *sitter.Node, fn func(*sitter.Node)) {
	fn(node)
	for i := 0; i < int(node.ChildCount()); i++ {
		tsWalkAll(node.Child(i), fn)
	}
}

// tsFindSymbol searches for a named symbol in the tree.
func tsFindSymbol(root *sitter.Node, src []byte, name string) *sitter.Node {
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)

		switch child.Type() {
		case "function_declaration":
			if tsNodeName(child, src) == name {
				return child
			}
		case "class_declaration":
			if tsNodeName(child, src) == name {
				return child
			}
			// Search methods within class
			if body := tsChildByType(child, "class_body"); body != nil {
				for j := 0; j < int(body.ChildCount()); j++ {
					member := body.Child(j)
					if member.Type() == "method_definition" && tsNodeName(member, src) == name {
						return member
					}
				}
			}
		case "interface_declaration", "type_alias_declaration":
			if tsNodeName(child, src) == name {
				return child
			}
		case "lexical_declaration", "variable_declaration":
			for j := 0; j < int(child.ChildCount()); j++ {
				decl := child.Child(j)
				if decl.Type() == "variable_declarator" && tsNodeName(decl, src) == name {
					return child
				}
			}
		case "export_statement":
			// Unwrap export and search inside
			for j := 0; j < int(child.ChildCount()); j++ {
				inner := child.Child(j)
				if inner.Type() == "function_declaration" && tsNodeName(inner, src) == name {
					return inner
				}
				if inner.Type() == "class_declaration" && tsNodeName(inner, src) == name {
					return inner
				}
				if inner.Type() == "interface_declaration" && tsNodeName(inner, src) == name {
					return inner
				}
				if inner.Type() == "type_alias_declaration" && tsNodeName(inner, src) == name {
					return inner
				}
				if inner.Type() == "lexical_declaration" || inner.Type() == "variable_declaration" {
					for k := 0; k < int(inner.ChildCount()); k++ {
						decl := inner.Child(k)
						if decl.Type() == "variable_declarator" && tsNodeName(decl, src) == name {
							return inner
						}
					}
				}
			}
		}
	}
	return nil
}
