package evaluator

import (
	"codys-monkey/ast"
	"codys-monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
