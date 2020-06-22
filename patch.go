package scim

import (
	"strings"

	filter "github.com/di-wu/scim-filter-parser"
)

const (
	// PatchOperationAdd is used to add a new attribute value to an existing resource.
	PatchOperationAdd = "add"
	// PatchOperationRemove removes the value at the target location specified by the required attribute "path".
	PatchOperationRemove = "remove"
	// PatchOperationReplace replaces the value at the target location specified by the "path".
	PatchOperationReplace = "replace"
)

var validOps = []string{PatchOperationAdd, PatchOperationRemove, PatchOperationReplace}

// PatchOperation represents a single PATCH operation.
type PatchOperation struct {
	// Op indicates the operation to perform and MAY be one of "add", "remove", or "replace".
	Op string
	// Path contains an attribute path describing the target of the operation. The "path" attribute is OPTIONAL for
	// "add" and "replace" and is REQUIRED for "remove" operations.
	Path string
	// Value specifies the value to be added or replaced.
	Value interface{}
}

// PatchRequest represents a resource PATCH request.
type PatchRequest struct {
	Schemas    []string
	Operations []PatchOperation
}

// GetPathFilter parses patch operation path to determine if it is a attribute filter.
// If it is, filter.Expression will be returned, nil otherwise.
func (p PatchOperation) GetPathFilter() *filter.AttributeExpression {
	parser := filter.NewParser(strings.NewReader(p.Path))
	pathFilter, err := parser.Parse()
	if err != nil {
		return nil
	}

	if attrFilter, ok := pathFilter.(filter.AttributeExpression); ok {
		return &attrFilter
	}
	return nil
}
