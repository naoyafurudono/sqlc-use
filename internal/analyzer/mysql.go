package analyzer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver" // required for parser
)

// MySQLAnalyzer analyzes MySQL queries
type MySQLAnalyzer struct {
	parser *parser.Parser
}

// NewMySQLAnalyzer creates a new MySQL analyzer
func NewMySQLAnalyzer() *MySQLAnalyzer {
	return &MySQLAnalyzer{
		parser: parser.New(),
	}
}

// Analyze extracts table effects from a MySQL query
func (a *MySQLAnalyzer) Analyze(queryName, sql string) (string, error) {
	// Parse the SQL
	stmtNodes, _, err := a.parser.Parse(sql, "", "")
	if err != nil {
		return "", fmt.Errorf("failed to parse SQL: %w", err)
	}

	operations := make(map[string]string)

	// Analyze each statement
	for _, stmtNode := range stmtNodes {
		visitor := &tableVisitor{
			operations: make(map[string]string),
		}

		// Determine operation type and visit nodes
		switch stmt := stmtNode.(type) {
		case *ast.SelectStmt:
			visitor.defaultOperation = "select"
			stmtNode.Accept(visitor)
		case *ast.SetOprStmt:
			// SetOprStmt handles UNION, EXCEPT, INTERSECT operations
			visitor.defaultOperation = "select"
			if stmt.SelectList != nil {
				for _, selectNode := range stmt.SelectList.Selects {
					selectNode.Accept(visitor)
				}
			}
		case *ast.InsertStmt:
			visitor.defaultOperation = "insert"
			stmtNode.Accept(visitor)
		case *ast.UpdateStmt:
			visitor.defaultOperation = "update"
			stmtNode.Accept(visitor)
		case *ast.DeleteStmt:
			visitor.defaultOperation = "delete"
			stmtNode.Accept(visitor)
		default:
			// Skip unsupported statement types
			continue
		}

		// Merge visitor results
		for table, operation := range visitor.operations {
			operations[table] = operation
		}
	}

	// Convert to effects string format
	return formatEffects(operations), nil
}

// tableVisitor visits AST nodes to extract table information
type tableVisitor struct {
	operations       map[string]string
	defaultOperation string
}

// Enter is called when entering a node
func (v *tableVisitor) Enter(node ast.Node) (ast.Node, bool) {
	switch n := node.(type) {
	case *ast.TableName:
		tableName := n.Name.String()
		if tableName != "" {
			// Use the default operation if not already set
			if _, exists := v.operations[tableName]; !exists {
				v.operations[tableName] = v.defaultOperation
			}
		}
	case *ast.TableSource:
		// Handle table sources in JOIN clauses
		if n.Source != nil {
			if tn, ok := n.Source.(*ast.TableName); ok {
				tableName := tn.Name.String()
				if tableName != "" && v.defaultOperation == "select" {
					v.operations[tableName] = "select"
				}
			}
		}
	}
	return node, false // false = continue traversal
}

// Leave is called when leaving a node
func (v *tableVisitor) Leave(node ast.Node) (ast.Node, bool) {
	return node, true // true = continue traversal
}

// formatEffects converts operation map to effects string format
func formatEffects(operations map[string]string) string {
	if len(operations) == 0 {
		return "{ }"
	}

	// Sort for consistent output
	var effects []string
	for table, operation := range operations {
		effects = append(effects, operation+"["+table+"]")
	}
	sort.Strings(effects)

	return "{ " + strings.Join(effects, " | ") + " }"
}
