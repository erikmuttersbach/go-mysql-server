package plan

import (
	"fmt"

	"github.com/src-d/go-mysql-server/sql"
)

type ShowStatus struct {
	pattern string
}

// NewShowStatus returns a new ShowStatus reference.
// config is a variables lookup table
// like is a "like pattern". If like is an empty string it will return all variables.
func NewShowStatus(like string) *ShowStatus {
	return &ShowStatus{
		pattern: like,
	}
}

// Resolved implements sql.Node interface. The function always returns true.
func (sv *ShowStatus) Resolved() bool {
	return true
}

// WithChildren implements the Node interface.
func (ss *ShowStatus) WithChildren(children ...sql.Node) (sql.Node, error) {
	if len(children) != 0 {
		return nil, sql.ErrInvalidChildrenNumber.New(ss, len(children), 0)
	}

	return ss, nil
}

// String implements the Stringer interface.
func (sv *ShowStatus) String() string {
	var like string
	if sv.pattern != "" {
		like = fmt.Sprintf(" LIKE '%s'", sv.pattern)
	}
	return fmt.Sprintf("SHOW STATUS%s", like)
}

// Schema returns a new Schema reference for "SHOW VARIABLES" query.
func (*ShowStatus) Schema() sql.Schema {
	return sql.Schema{
		&sql.Column{Name: "Variable_name", Type: sql.Text, Nullable: false},
		&sql.Column{Name: "Value", Type: sql.Text, Nullable: true},
	}
}

// Children implements sql.Node interface. The function always returns nil.
func (*ShowStatus) Children() []sql.Node { return nil }

// RowIter implements the sql.Node interface.
// The function returns an iterator for filtered variables (based on like pattern)
func (sv *ShowStatus) RowIter(ctx *sql.Context) (sql.RowIter, error) {
	var rows []sql.Row
	return sql.RowsToRowIter(rows...), nil
}
