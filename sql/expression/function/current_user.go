package function

import (
	"github.com/src-d/go-mysql-server/sql"
)

type CurrentUser struct {
}

func NewCurrentUser() sql.Expression {
	return &CurrentUser{}
}

// Type implements the sql.Expression interface.
func (*CurrentUser) Type() sql.Type { return sql.Text }

func (*CurrentUser) String() string { return "CURRENT_USER()" }

// IsNullable implements the sql.Expression interface.
func (*CurrentUser) IsNullable() bool { return false }

// Resolved implements the sql.Expression interface.
func (*CurrentUser) Resolved() bool { return true }

// Children implements the sql.Expression interface.
func (*CurrentUser) Children() []sql.Expression { return nil }

// Eval implements the sql.Expression interface.
func (n *CurrentUser) Eval(ctx *sql.Context, _ sql.Row) (interface{}, error) {
	return ctx.Session.Client().User+"@localhost", nil
}

// WithChildren implements the Expression interface.
func (n *CurrentUser) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	if len(children) != 0 {
		return nil, sql.ErrInvalidChildrenNumber.New(n, len(children), 0)
	}
	return n, nil
}

