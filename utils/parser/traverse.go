package parser

import (
	"go/ast"
)

// TraverseAST traverse the given AST node and call the given function for each node
func TraverseAST(node *ast.File, fn func(ast.Node) bool) {
	// Traverse the AST to find the struct and field
	ast.Inspect(
		node, fn,
	)
}
