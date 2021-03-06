package gorma

import (
	"strings"

	"github.com/jinzhu/inflection"
)

// LeftNamePlural returns the pluralized version of
// the "owner" of the m2m relationship.
func (m *ManyToManyDefinition) LeftNamePlural() string {
	return inflection.Plural(m.Left.ModelName)
}

// RightNamePlural returns the pluralized version
// of the "child" of the m2m relationship.
func (m *ManyToManyDefinition) RightNamePlural() string {
	return inflection.Plural(m.Right.ModelName)
}

// LeftName returns the name of the "owner" of the
// m2m relationship.
func (m *ManyToManyDefinition) LeftName() string {
	return m.Left.ModelName
}

// RightName returns the name of the "child" of the
// m2m relationship.
func (m *ManyToManyDefinition) RightName() string {
	return m.Right.ModelName
}

// LowerLeftName returns the lower case name of the "owner" of the
// m2m relationship.
func (m *ManyToManyDefinition) LowerLeftName() string {
	return strings.ToLower(m.Left.ModelName)
}

// LowerRightName returns the name of the "child" of the
// m2m relationship.
func (m *ManyToManyDefinition) LowerRightName() string {
	return strings.ToLower(m.Right.ModelName)
}
